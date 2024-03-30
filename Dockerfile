FROM golang:1.22 as builder
ADD ./app /app
WORKDIR /app
RUN go build -o api /app/*.go

FROM debian:stable-slim
WORKDIR /app
COPY --from=builder /app/api .
CMD ["./api"]