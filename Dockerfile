FROM golang:alpine as builder

RUN apk add git
RUN go get -u github.com/golang/dep/cmd/dep

ADD . /go/src/restgo_framework
WORKDIR /go/src/restgo_framework


RUN dep init -v
RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o main

FROM scratch
COPY --from=builder /go/src/restgo_framework/main /main
COPY --from=builder /go/src/restgo_framework/.env /.env
WORKDIR /
CMD ["./main"]
