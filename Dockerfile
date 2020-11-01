FROM golang:1.15-alpine as builder

ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

RUN mkdir /app
COPY . /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -installsuffix cgo -o main .

FROM alpine:latest AS runtime
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080

CMD ["./main"]