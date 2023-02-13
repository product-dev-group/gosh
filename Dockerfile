FROM ubuntu:latest

RUN apt-get update && apt-get install -y vim git

RUN apt-get install software-properties-common -y

RUN add-apt-repository ppa:longsleep/golang-backports &&\
    apt install golang-go -y

RUN mkdir -p /opt/gosh

WORKDIR /opt/gosh

COPY . .

CMD ["/bin/bash"]