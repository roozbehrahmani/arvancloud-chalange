FROM golang:1.17

COPY ./discount-service /go/src/github.com/roozbehrahmani/arvancloud-chalange
WORKDIR /go/src/github.com/roozbehrahmani/arvancloud-chalange
ENV BIN=/go/src/github.com/roozbehrahmani/arvancloud-chalange/bin
ENV GOBIN=/go/src/github.com/roozbehrahmani/arvancloud-chalange/bin
# RUN go mod tidy
RUN go mod vendor
RUN go install -mod=vendor -v -ldflags "-s" ./cmd/app/...

CMD ["/go/src/github.com/roozbehrahmani/arvancloud-chalange/bin/app"]

EXPOSE 8080