FROM golang:latest 
RUN mkdir /marketbot 
ADD . /marketbot/ 
WORKDIR /marketbot
RUN go build -o tbot cmd/tbot/main.go 
CMD ["/marketbot/tbot"]