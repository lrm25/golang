FROM ubuntu

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata
RUN apt-get install -y apt-utils golang git
RUN cd /home && git clone https://github.com/lrm25/golang
CMD go run /home/golang/pickanumber/pickanumber.go
