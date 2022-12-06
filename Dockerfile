FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . ./

RUN go build -o /promotion-app

EXPOSE 8080

CMD [ "/promotion-app"]