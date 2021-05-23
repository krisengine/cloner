FROM golang:1.14 as builder
LABEL maintainer="Roman S"
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/tidwall/gjson
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

FROM alpine:3.10
RUN apk --no-cache add ca-certificates git
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x ./main
EXPOSE 8000
CMD ["/app/main"]