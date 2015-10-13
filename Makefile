
all: compile

test:
	go test -bench .

compile:
	go build bin/catgr.go
	go build bin/demoreport.go

arch:
	GOOS="darwin"  GOARCH="amd64" go build -o catgr-darwin64 bin/catgr.go
	GOOS="linux"   GOARCH="amd64" go build -o catgr-linux64 bin/catgr.go
	GOOS="windows" GOARCH="amd64" go build -o catgr-win64.exe bin/catgr.go

