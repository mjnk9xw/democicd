FROM golang:alpine as BUILDER
WORKDIR /go/src/app
COPY . .
RUN go mod vendor
RUN GOOS=linux go build -buildvcs=false -o main
FROM alpine
EXPOSE 80
WORKDIR /app
COPY --from=BUILDER /go/src/app/main /app
CMD /app/main