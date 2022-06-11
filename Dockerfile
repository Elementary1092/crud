FROM golang:latest

RUN go install github.com/Elementary1092/crud@latest
RUN git clone --depth 1 https://github.com/Elementary1092/crud

ENV GO111MODULE=on
ENV APP_DIR /crud/build

RUN go build -o $APP_DIR /crud

EXPOSE 8081
CMD ["/crud/build"]