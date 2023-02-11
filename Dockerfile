FROM ubuntu:latest

RUN apt-get update && apt-get install -y vim git && apt install sudo

RUN sudo apt-get install software-properties-common -y

RUN sudo apt update &&\
    sudo add-apt-repository ppa:longsleep/golang-backports &&\
    sudo apt install golang-go -y

RUN mkdir -p root/app

WORKDIR root/app

COPY . .

CMD ["/bin/bash"]