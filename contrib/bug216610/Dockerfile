FROM debian:latest

# A directory to share files via.
RUN mkdir /shared

RUN apt-get update
RUN apt-get install -y gcc-arm-linux-gnueabi binutils-arm-linux-gnueabi
RUN apt-get install -y gcc-aarch64-linux-gnu binutils-aarch64-linux-gnu

# create a builder user
RUN echo "builder:x:1000:1000:,,,:/home/builder:/bin/bash" >> /etc/passwd
RUN echo "builder:*:19289:0:99999:7:::" >> /etc/shadow
RUN mkdir -p /home/builder && chown builder.bin /home/builder
