FROM golang:1.26

ARG SSH_PRIVATE_KEY=""
WORKDIR /workspace
RUN apt-get update && apt-get install -y sudo
## operator-sdk settings
RUN echo "export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac)" >> ~/.bashrc
RUN echo "export OS=$(uname | awk '{print tolower($0)}')" >> ~/.bashrc
RUN echo "export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/v1.42.2" >> ~/.bashrc

## copying file because source command not working in Dockerfile
COPY reload_and_download.sh reload_and_download.sh
RUN cat /workspace/reload_and_download.sh
RUN sudo chmod +x /workspace/reload_and_download.sh
RUN /workspace/reload_and_download.sh

# RUN gpg --keyserver keyserver.ubuntu.com --recv-keys 052996E2A20B5C7E
# RUN curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt && curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt.asc


# RUN chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk

## SSH settings
RUN mkdir ~/.ssh && chmod -R 700 ~/.ssh

RUN echo "${SSH_PRIVATE_KEY}" >> ~/.ssh/id_rsa

RUN chmod 0400 ~/.ssh/id_rsa && echo "StrictHostKeyChecking no" > ~/.ssh/config
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts



RUN git clone git@github.com:jmzk96/RedOps.git 