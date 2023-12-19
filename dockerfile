# Stage 1: Build the application
FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download

# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -o main .

# Stage 2: Create the final container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app .

EXPOSE 7777
#Default cant be overriden
ENTRYPOINT [ "./main" ]
#Comandline args which can be overriden
CMD ["./main"]