FROM golang:latest 
RUN mkdir /mbserver 
ADD ../.. /mbserver/ 
WORKDIR /mbserver
RUN go build -o tbot cmd/server/main.go 
CMD ["/mbserver/tbot"]