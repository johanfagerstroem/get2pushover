FROM golang:1 as builder
COPY . /get2pushover
RUN cd /get2pushover && CGO_ENABLED=0 make

FROM alpine
WORKDIR /root
COPY --from=builder /get2pushover/build/bin/get2pushover .
CMD ["/root/get2pushover"]