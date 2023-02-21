FROM golang:1.20-bullseye

WORKDIR /go-rest-template

RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y \
        procps \
        vim \
        less \
        telnet \
        curl \
        net-tools\
        upx-ucl

COPY go.mod .
COPY go.sum .

COPY . .

RUN make download
RUN make build

EXPOSE 8080

CMD ["./grt"]