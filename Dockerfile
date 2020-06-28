FROM golang
LABEL maintainer="extrawurst"
COPY main.go go.mod go.sum /ip2loc/
RUN cd /ip2loc && go build -o ip2location main.go
WORKDIR /ip2loc
ENTRYPOINT ["./ip2location"]