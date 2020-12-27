FROM golang:1.15 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin/app ./
RUN ldd /go/bin/app | tr -s '[:blank:]' '\n' | grep '^/' | \
    xargs -I % sh -c 'mkdir -p $(dirname deps%); cp % deps%;'

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/deps /
COPY --from=builder /go/bin/app /go/bin/app

EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]
