FROM golang:1.22

WORKDIR /github.com/Elvilius/user-events-audit-hub

COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/service/" -command="./service"