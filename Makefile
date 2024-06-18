.PHONY: default

default:
	rm -fv bin/pingmullvad-*
	GOOS=darwin GOARCH=arm64 go build -o bin/pingmullvad-arm64-0.2.0 -ldflags "-X main.ver=0.2.0 -X 'main.build=`date +%Y%m%d%H%M%S%3N`'" main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/pingmullvad-amd64-0.2.0 -ldflags "-X main.ver=0.2.0 -X 'main.build=`date +%Y%m%d%H%M%S%3N`'" main.go
