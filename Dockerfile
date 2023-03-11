FROM golang:1.17-alpine

WORKDIR /app

COPY ./server /app/
RUN go mod download
RUN go build -o /one-time-secret

EXPOSE 8081

CMD [ "./one-time-secret" ]