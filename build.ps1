$sourcecode = "main.go"
$target = "build/health-checker"
$env:GOOS = 'windows'; $env:GOARCH = 'amd64';               go build -o "$($target)-win64.exe" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'amd64';               go build -o "$($target)-linux64" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'arm'; $env:GOARM=5;   go build -o "$($target)-raspi32" $sourcecode
$env:GOOS = 'darwin';  $env:GOARCH = 'amd64';               go build -o "$($target)-macos64.macos" $sourcecode