FROM golang:rc-buster

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build . && rm -rf go.* internal *.go pkg && go clean

CMD ["./pr0.bot"]
