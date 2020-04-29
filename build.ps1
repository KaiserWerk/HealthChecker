$sourcecode = "main.go"
$target = "build/WebsiteHealthChecker"
$env:GOOS = 'windows'; $env:GOARCH = 'amd64';               go build -o "$($target)_win64.exe" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'amd64';               go build -o "$($target)_linux64" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'arm'; $env:GOARM=5;   go build -o "$($target)_raspi32" $sourcecode
$env:GOOS = 'darwin';  $env:GOARCH = 'amd64';               go build -o "$($target)_macos64.macos" $sourcecode