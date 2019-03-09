FROM golang:alpine as builder
RUN mkdir -p /build
ADD ./*.go /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -o api .
FROM scratch
COPY --from=builder /build/api /go/bin/api
EXPOSE 60080
