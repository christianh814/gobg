# Build the App
FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o /app/gobg

# Build the Image
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/gobg /app/gobg

COPY --from=build /app/html /app/html

EXPOSE 8080

USER 1001

CMD ["/app/gobg"]