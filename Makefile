build:
	env GOOS=linux CGO_ENABLED=0 GO111MODULE=on /usr/local/go/bin/go build -mod=vendor -o builds/parse_jwt_token main.go

build-win:
	env GOOS=windows CGO_ENABLED=0 GO111MODULE=on /usr/local/go/bin/go build -mod=vendor -o builds/parse_jwt_token.exe main.go