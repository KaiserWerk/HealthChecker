# HealthChecker

A simple host health checking tool. To get started, clone the repository and build for
your required system. If you're on Windows, you can use the ``build.ps1`` PowerShell script
to cross-compile for Windows 64bit, Linux 64bit, RaspberryPi (ARM) 32bit 
(suitable for most RaspberryPi's) and MacOS 64bit.

To run, you will need two files, `config.json` and `urls.json`.

The `config.json` file contains the [PushOver](https://pushover.net/) User key and 
API key, e.g.:

    {
      "userkey": "<my-user-key>",
      "apikey": "<my-api-key>"
    }

The `urls.json` file contains a list of URLs to check, e.g.:

    [
      "https://github.com",
      "http://127.0.0.1:4001/ping"
    ]
    
When requests to the PushOver API fail, a log entry will be written to the 
`response_errors.log` file.