FROM golang:rc-buster

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build . && rm -rf go.* internal *.go pkg && go clean

CMD ["./pr0.bot"]