FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg pkg
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o /processor main.go

EXPOSE 8080

CMD ["/processor"]
