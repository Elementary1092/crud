FROM golang:latest

ENV GO111MODULE=on

RUN git clone --depth 1 https://github.com/Elementary1092/crud

WORKDIR /crud

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o build/ cmd/main.go

EXPOSE 8081
CMD ["build/main"]