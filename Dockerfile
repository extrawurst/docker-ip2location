FROM golang:alpine
LABEL maintainer="extrawurst"
COPY data/* /ip2loc/data/
COPY main.go go.mod go.sum /ip2loc/
RUN cd /ip2loc && go build -o ip2location main.go
WORKDIR /ip2loc
ENTRYPOINT ["./ip2location"]