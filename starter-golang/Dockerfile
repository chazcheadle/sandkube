# Stage 1
# Build stage to compile Go binary
FROM golang:1-alpine as build
ADD ./src/main.go /go/src
WORKDIR /go/src
RUN go build main.go

# Stage 2
# Build stage to run minimal OS and binary
FROM alpine:latest
WORKDIR /app
# Copy executable to new image
COPY --from=build /go/src/main /app/main
EXPOSE 8080
ENTRYPOINT ["./main"]