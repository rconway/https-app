FROM golang:1.12-alpine as builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -v

FROM alpine:latest
RUN apk --no-cache add ca-certificates
EXPOSE 8000/tcp
COPY --from=builder /app/webappgo /webappgo
ENTRYPOINT ["/webappgo"]
CMD []