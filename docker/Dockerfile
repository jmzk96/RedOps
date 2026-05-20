FROM golang:1.26

RUN export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac) &&\\
export OS=$(uname | awk '{print tolower($0)}')

RUN export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/v1.42.2 &&\\
curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}

RUN gpg --keyserver keyserver.ubuntu.com --recv-keys 052996E2A20B5C7E

RUN grep operator-sdk_${OS}_${ARCH} checksums.txt | sha256sum -c -

RUN chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk

WORKDIR /workspace

RUN git clone git@github.com:jmzk96/RedOps.git 

WORKDIR /workspace/RedOps