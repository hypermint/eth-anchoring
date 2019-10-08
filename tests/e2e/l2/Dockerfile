FROM bluele/hypermint:0.4.4
ENV TM_PARAMS "consensus.timeout_commit=100ms,rpc.max_body_bytes=10000000,mempool.max_msg_bytes=10000000"
ENV HMD_HOME /root/.hmd
ENV HMCLI_HOME /root/.hmcli
ENV MNEMONIC "hazard decade force acid harbor praise wagon dragon they increase build drink"
ENV HDW_PATH "m/44'/60'/0'/0"
ENV HDW_VALIDATOR_IDX '0'

COPY ./setup.sh /setup.sh
RUN sh /setup.sh
