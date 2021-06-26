# This Dockerfile generates an 
FROM golang:1.16.5 as builder
WORKDIR /yadb.d
COPY ./ /yadb.d
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-w -s" -o yadb cmd/*.go

FROM alpine
WORKDIR /
COPY --from=builder /yadb.d/yadb /
ENTRYPOINT [ "/yadb" ]
