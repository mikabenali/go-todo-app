FROM golang:1.21.7

WORKDIR /app

COPY /app .

RUN go get
RUN go build -o bin .


ENTRYPOINT [ "/app/bin" ]
