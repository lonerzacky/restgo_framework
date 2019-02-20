build:
	docker build -t go-sample:latest .
stop:
	docker stop myGoApp
rm-container:
	docker rm myGoApp
rm-image:
	docker rmi go-sample:latest
run:
	docker run -p 1000:1000 -v C:\Users\wolfz\go\src\restgo_framework\.env:/.env --name myGoApp go-sample
gobuild:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o main