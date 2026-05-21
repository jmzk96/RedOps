#!/bin/bash

source ~/.bashrc

# download operator-sdk binary
curl -LO "${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}"

# confirm the binary is valid
gpg --keyserver keyserver.ubuntu.com --recv-keys 052996E2A20B5C7E
curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt
curl -LO ${OPERATOR_SDK_DL_URL}/checksums.txt.asc
gpg -u "Operator SDK (release) <cncf-operator-sdk@cncf.io>" --verify checksums.txt.asc

chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk