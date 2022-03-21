FROM golang:1.18

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

WORKDIR /usr/src/app/App
RUN go build -v -o /usr/local/bin/app
WORKDIR /usr/local/bin/

EXPOSE 8080 27017

CMD ["app"]




