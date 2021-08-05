FROM golang:latest
RUN mkdir /go/src/charts_server
WORKDIR /go/src/charts_server
ADD . /go/src/charts_server

RUN go get github.com/go-sql-driver/mysql;
RUN go get -u github.com/gorilla/mux

ARG DOCKER_UID=1000
ARG DOCKER_USER=docker
ARG DOCKER_PASSWORD=docker
RUN useradd -m --uid ${DOCKER_UID} --groups sudo ${DOCKER_USER} \
  && echo ${DOCKER_USER}:${DOCKER_PASSWORD} | chpasswd
USER ${DOCKER_USER}

