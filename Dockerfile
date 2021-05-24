FROM golang:1.14 as builder
LABEL maintainer="Roman S"
WORKDIR /app
COPY . .
RUN ./build.sh

FROM alpine:3.10
RUN apk --no-cache add ca-certificates git
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["/app/main"]