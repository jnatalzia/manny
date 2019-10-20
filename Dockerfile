FROM golang:1.12

ENV PROJECTROOT=/go/src/github.com/jnatalzia/go-boilerplate

RUN mkdir -p /tmp
RUN mkdir -p $PROJECTROOT

WORKDIR $PROJECTROOT
COPY ./main.go .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build

CMD ./go-boilerplate -t $AUTH_KEY