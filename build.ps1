$sourcecode = "main.go"
$target = "build/health-checker"
$env:GOOS = 'windows'; $env:GOARCH = 'amd64';               go build -o "$($target)-win64.exe"     -ldflags "-s -w" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'amd64';               go build -o "$($target)-linux64"       -ldflags "-s -w" $sourcecode
$env:GOOS = 'linux';   $env:GOARCH = 'arm'; $env:GOARM=5;   go build -o "$($target)-raspi32"       -ldflags "-s -w" $sourcecode
$env:GOOS = 'darwin';  $env:GOARCH = 'amd64';               go build -o "$($target)-macos64.macos" -ldflags "-s -w" $sourcecode
