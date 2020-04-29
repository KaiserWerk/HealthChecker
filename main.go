package main

import (
	"encoding/json"
	"fmt"
	"github.com/gregdel/pushover"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Parameters struct {
	UserKey	string	`json:"userkey"`
	ApiKey	string	`json:"apikey"`
}

type CheckUrl struct {
	Url       			string
	IsOffline 			bool
	IsPreviouslyOffline bool
	IsSent    			bool
}

var (
	Urls []CheckUrl
	client = http.Client{
		Timeout: 6 * time.Second,
	}
	parameters Parameters
	App        *pushover.Pushover
)

func main() {
	// read and set parameters
	params, err := getParameters()
	if err != nil {
		panic(err.Error())
	}
	parameters = params

	// create pushover app "instance"
	App = pushover.New(parameters.ApiKey)

	urlList, err := readUrls()
	if err != nil {
		panic(err.Error())
	}
	// fill slice of CheckUrls
	for _, v := range urlList {
		Urls = append(Urls, CheckUrl{Url: v})
	}


	ticker := time.NewTicker(1 * time.Minute)

	for _ = range ticker.C {
		checkAllUrls()
		var str strings.Builder
		for k, v := range Urls {
			if v.IsOffline == true {
				if v.IsSent == false {
					// message zusammenbauen
					if v.IsPreviouslyOffline == true {
						str.WriteString(v.Url + " is still offline!\n")
					} else {
						str.WriteString(v.Url + " is offline!\n")
					}
					Urls[k].IsSent = true
				}
			} else {
				if v.IsPreviouslyOffline == true {
					Urls[k].IsSent = true
					str.WriteString(v.Url + " is back online!")
				}
			}
		}

		// send message, if message not empty
		if str.Len() > 0 {
			sendNotification(str.String())
			str.Reset()
		}
	}

}

func readUrls() ([]string, error) {
	cont, err := ioutil.ReadFile("urls.json")
	if err != nil {
		return nil, err
	}

	var urls []string

	err = json.Unmarshal(cont, &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func getParameters() (Parameters, error) {
	cont, err := ioutil.ReadFile("config.json")
	if err != nil {
		return Parameters{}, err
	}
	var params Parameters
	err = json.Unmarshal(cont, &params)
	if err != nil {
		return Parameters{}, err
	}

	return params, nil
}

func sendNotification(messageText string) {
	recipient := pushover.NewRecipient(parameters.UserKey)
	// Create the message to send
	message := &pushover.Message{
		Priority:   pushover.PriorityNormal,
		Message:    messageText,
		Title:      "Site(s) offline",
		Sound:      pushover.SoundBike,
	}

	// Send the message to the recipient
	response, err := App.SendMessage(message, recipient)
	if err != nil {
		fmt.Println("could not send notification: " + err.Error())
	}

	if response != nil && response.Status != 1 {
		jsonBytes, err := json.Marshal(response)
		if err != nil {
			fmt.Println("could not parse response to json")
		}
		jsonValue := string(jsonBytes) + "\n"
		filename := "response_errors.log"
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0640)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer f.Close()

		if _, err = f.WriteString(jsonValue); err != nil {
			fmt.Println("could not write response to log file")
			return
		}
	}
}

func checkAllUrls() {
	for k, v := range Urls {
		if !checkUrl(v) {
			Urls[k].IsPreviouslyOffline = Urls[k].IsOffline
			Urls[k].IsOffline = true
			fmt.Println(v.Url + " is offline!")
		} else {
			Urls[k].IsPreviouslyOffline = Urls[k].IsOffline
			Urls[k].IsOffline = false
			fmt.Println(v.Url + " is online!")
		}
	}
	fmt.Println("")
}

func checkUrl(checkUrl CheckUrl) bool {
	response, errors := client.Get(checkUrl.Url)

	if errors != nil {
		_, netErrors := client.Get("https://www.google.com")

		if netErrors != nil {
			fmt.Println("no internet connection; exiting...")
			os.Exit(0)
		}

		return false
	}

	if response.StatusCode < 400 {
		return true
	}

	return false
}