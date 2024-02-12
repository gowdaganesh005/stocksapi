FROM golang:1.12.0-alpine3.9
RUN mkdir /stocksapi
ADD . /stocksapi
WORKDIR /stocksapi
RUN go build -o main .
CMD ["/stockapi/main"]