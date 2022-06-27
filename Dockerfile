FROM golang:latest 
RUN mkdir /marketbot 
ADD . /marketbot/ 
WORKDIR /marketbot

RUN go build -o tbot cmd/tbot/main.go 
RUN go build -o server cmd/server/main.go 
RUN go build -o myinvest cmd/myinvest/main.go 

CMD ["/marketbot/scripts/start.sh"]