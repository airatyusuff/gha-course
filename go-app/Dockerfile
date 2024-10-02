FROM golang:1.23.1-alpine3.19 AS build
WORKDIR /app
COPY go-app/ .
COPY go-app/go.mod ./
RUN go mod download && go mod verify
RUN go build -o main .

# Final stage
FROM scratch
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]