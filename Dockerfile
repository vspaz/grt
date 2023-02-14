FROM golang:1.20-bullseye

WORKDIR /go-rest-template

RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y \
        procps \
        vim \
        less \
        telnet \
        curl \
        net-tools

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o grt

EXPOSE 8080

CMD ["./grt"]