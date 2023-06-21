FROM ubuntu:22.04

RUN apt-get update && \
    apt-get install -y curl tar

RUN curl --location --output /tmp/kaf.tar.gz https://github.com/birdayz/kaf/releases/download/v0.2.3/kaf_0.2.3_Linux_x86_64.tar.gz && \
    tar -zxvf /tmp/kaf.tar.gz && \
    mv kaf /usr/local/bin && \
    kaf --version

COPY config.yaml /root/.kaf/config
