FROM golang:1.5
ADD . $GOPATH/src/github.com/e-r-w/flying-squid
RUN go get -u github.com/aws/aws-sdk-go/...
RUN go get -u github.com/go-martini/martini
RUN go get -u github.com/martini-contrib/render
RUN go install github.com/e-r-w/flying-squid
ENTRYPOINT $GOPATH/bin/flying-squid
