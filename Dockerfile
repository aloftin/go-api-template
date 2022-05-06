FROM golang:1.18-alpine AS builder
RUN apk --no-cache add git
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main .

FROM golang:1.18-alpine
RUN apk --no-cache add openssl
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
CMD ["/app/main"]