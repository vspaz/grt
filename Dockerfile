FROM golang:1.20-bullseye

WORKDIR /go-rest-template

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o grt

EXPOSE 8080

CMD ["./grt"]