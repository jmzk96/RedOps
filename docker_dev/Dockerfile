FROM golang:1.26

ARG SSH_PRIVATE_KEY=""

RUN export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac) && \
export OS=$(uname | awk '{print tolower($0)}')

RUN export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/v1.42.2 && \
curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}

RUN apt-get update && apt-get install -y sudo

RUN chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk

# SSH settings
RUN mkdir ~/.ssh && chmod -R 700 ~/.ssh

RUN echo "${SSH_PRIVATE_KEY}" >> ~/.ssh/id_rsa
RUN echo ${SSH_PRIVATE_KEY}

RUN chmod 0400 ~/.ssh/id_rsa && echo "StrictHostKeyChecking no" > ~/.ssh/config
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts


WORKDIR /workspace
RUN git clone git@github.com:jmzk96/RedOps.git 
