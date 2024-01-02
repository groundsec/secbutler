FROM ubuntu:24.04

ARG DEBIAN_FRONTEND=noninteractive
RUN apt-get update
RUN apt-get install -y \
  ca-certificates \
  gnupg \
  sudo \
  cmake \
  gcc \
  libssl-dev \
  automake \
  xterm \
  proxychains \
  vim \
  git \
  python3 \
  python3-pip \
  wget \
  curl \
  libpcap-dev \
  pipx
RUN mkdir -p /etc/apt/keyrings
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_20.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list
RUN apt update && apt install -y nodejs
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
RUN wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz 
RUN rm -rf /usr/local/go
RUN tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
RUN echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
RUN echo 'export PATH="$PATH:/root/.local/bin"' >> ~/.bashrc
