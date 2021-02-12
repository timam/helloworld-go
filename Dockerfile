FROM golang:1.14 AS build

WORKDIR /go/src/
COPY . /go/src/

WORKDIR /go/src
RUN CGO_ENABLED=0 go build -o /bin/helloworld-go

FROM scratch
COPY --from=build /bin/helloworld-go /bin/helloworld-go
ENTRYPOINT ["/bin/helloworld-go"]