FROM golang 

WORKDIR /go/src/github.com/vlasove/grpcproject
COPY cmd/server cmd/server 
COPY pkg pkg
WORKDIR /go/src/github.com/vlasove/grpcproject/cmd/server

RUN go get -d -v
RUN go build 

CMD  ./server
EXPOSE 8081