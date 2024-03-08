FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go .
COPY data/ /app/data/
COPY views/ /app/views/
COPY static/ /app/static/

RUN CGO_ENABLED=0 GOOS=linux go build -o /medotlink

EXPOSE 5890

CMD ["/medotlink"]