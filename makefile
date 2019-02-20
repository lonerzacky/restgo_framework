REMOTE_USERNAME=root
REMOTE_IP=10.140.14.90
REMOTE_DESTINATION=/home/golang/src/restgo_framework

go-init:
	go get -u github.com/golang/dep/cmd/dep && dep init -v && ep ensure -v
build:
	docker build -t go-sample:latest .
stop:
	docker stop myGoApp
rm-container:
	docker rm myGoApp
rm-image:
	docker rmi go-sample:latest
clean:
	docker stop myGoApp && docker rm myGoApp && docker rmi go-sample:latest
run:
	docker run -p 1000:1000 -v C:\Users\wolfz\go\src\restgo_framework\.env:/.env --name myGoApp go-sample
gobuild:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o main
transfer:
	scp .env main makefile ${REMOTE_USERNAME}@${REMOTE_IP}:${REMOTE_DESTINATION}
docker-scratch:
	@echo -e "FROM scratch\nADD . .\nCMD ["./main"]" > Dockerfile