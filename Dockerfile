FROM golang:1.13 AS build

WORKDIR /go/src/
COPY src/ /go/src/

WORKDIR /go/src/hello/
RUN CGO_ENABLED=0 go build -o /bin/hello

FROM scratch
COPY --from=build /bin/hello /bin/hello
ENTRYPOINT ["/bin/hello"]







