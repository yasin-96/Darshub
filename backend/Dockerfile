FROM golang:1.18

WORKDIR /usr/src/app

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

CMD ["go run main.go"]