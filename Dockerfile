FROM golang:1.17

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /rockstock-go-api

EXPOSE 12345

CMD [ "/rockstock-go-api" ]