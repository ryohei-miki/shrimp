FROM golang:1.14

ENV GOBIN=/go/bin
ENV GO111MODULE=on
ENV GOPATH=

WORKDIR /go
COPY ./src ./src

COPY go.* ./
RUN go mod download

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $GOBIN

CMD ["go", "run", "src/main.go"]