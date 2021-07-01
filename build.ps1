$env:CGO_ENABLED=1;
$env:GOOS="windows";
$env:GOARCH=386;

go build -o test.dll main.go