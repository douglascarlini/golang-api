FROM golang:1.20-alpine

RUN mkdir /app

COPY src /app

WORKDIR /app

COPY swagger.json .
RUN go build -o server .

EXPOSE 80

CMD [ "/app/server" ]