FROM golang:1.23.1-alpine3.19 AS build
WORKDIR /app
COPY go-app/ .
COPY go-app/go.mod ./
RUN go mod download && go mod verify
RUN go build -o main .

# Migrations stage
WORKDIR /app
COPY go-app/migrations /app/migrations

# Final stage
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/migrations /app/migrations
EXPOSE 8080
CMD ["./main"]