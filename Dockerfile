FROM golang:latest
WORKDIR /app
COPY main.go .
RUN go get -d -v github.com/gorilla/mux
RUN go build main.go
CMD ["./main"]