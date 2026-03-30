FROM golang:1.25.0

WORKDIR /api

# Download modules (such as gin)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /api/cmd/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o album-app

# Run
CMD ["./album-app"]
