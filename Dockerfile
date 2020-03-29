FROM golang:latest

RUN mkdir $GOPATH/src/coronabot

COPY . $GOPATH/src/coronabot

WORKDIR $GOPATH/src/coronabot

RUN go get
RUN go build

ENTRYPOINT $GOPATH/src/coronabot/coronabot