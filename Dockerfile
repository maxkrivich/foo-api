FROM golang:1.17 as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o /go/bin/app .

FROM alpine:3.12
EXPOSE 8080
COPY --from=build /go/bin/app /
CMD ["/app"]
