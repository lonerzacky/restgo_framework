build:
	docker build -t go-sample:latest .
stop:
	docker stop myFlask
rm-container:
	docker rm myFlask
rm-image:
	docker rmi flask-sample:latest
run:
	docker run -p 7000:7000 -v C:\Users\wolfz\go\src\restgo_framework\.env:/.env --name myGoApp go-sample
