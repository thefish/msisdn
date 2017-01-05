FROM golang:1.7

RUN mkdir -p /go/src/msisdn
WORKDIR /go/src/msisdn
EXPOSE 8000

CMD ["./msisdn"]

COPY . /go/src/msisdn
RUN go build

