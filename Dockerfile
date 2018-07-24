FROM golang:1.8 as builder
WORKDIR /go/src/app
COPY . . 
RUN go get -d -v ./...
RUN go install -v ./...

FROM golang:1.8
COPY --from=builder /go/bin/app /go/bin
CMD /go/bin/app