FROM golang:latest


RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app/
COPY ./ /usr/src/app/

RUN go mod download

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]