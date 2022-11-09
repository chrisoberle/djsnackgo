FROM golang:1.18

LABEL base.name="djsnack.com"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main"]
