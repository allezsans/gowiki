FROM golang:1.12 as builder

WORKDIR /go/src/github.com/allezsans/gowiki
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o gowiki

FROM alpine

COPY --from=builder /go/src/github.com/allezsans/gowiki/gowiki /gowiki
CMD ["/gowiki"]