FROM golang:latest

RUN git clone --depth 1 https://github.com/Elementary1092/crud.git

ENV GO111MODULE=on
ENV APP_DIR /crud/build

RUN go build -o $APP_DIR /crud

WORKDIR /crud
EXPOSE 8081
CMD ["/crud"]