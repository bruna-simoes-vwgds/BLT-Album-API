FROM golang:1.25.0

WORKDIR /

# Download modules (such as gin)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /cmd/api

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o album-app

# Run
CMD ["./album-app"]
