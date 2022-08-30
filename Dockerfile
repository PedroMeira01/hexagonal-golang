FROM golang:1.16

WORKDIR /go/src
ENV GOPATH=/home/www-data/go
ENV PATH="/home/www-data/go/bin:${PATH}"
ENV GOROOT=/usr/local/go #gosetup

RUN go install github.com/spf13/cobra/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0

RUN apt-get update && apt-get install sqlite3 -y

RUN usermod -u 1000 www-data

USER www-data

CMD ["tail", "-f", "/dev/null"]