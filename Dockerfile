
FROM golang:1.16 as dev

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o url-shortner ./src/

FROM dev as debug

RUN go get  github.com/go-delve/delve/cmd/dlv@v1.6.0
RUN go get github.com/cespare/reflex@v0.3.1

RUN go get -d -v ./...


CMD reflex -R "__debug_bin" -s -- sh -c "dlv debug --headless --continue --accept-multiclient --listen :40000 --api-version=2 --log ./src/"


