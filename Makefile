
build:
	export CGO_ENABLED=0
	go get github.com/tools/godep
	godep go build -tags netgo -installsuffix s3t -v -a -ldflags '-s' -o bin/s3t
