FROM offscale/golang-builder-alpine3.8 as stage0

ADD bootstrap.sh /tmp/bootstrap.sh
RUN /tmp/bootstrap.sh

FROM scratch as lachesis_base

ENV node_num=0
ENV node_addr='127.0.0.1'

EXPOSE 1338
EXPOSE 1339
EXPOSE 8000
EXPOSE 12000

# cp -r /etc/ssl/certs certs, then add to your `docker build`: `--build-arg ca_certificates=certs`
ARG ca_certificates
ADD "$ca_certificates" /etc/ssl/certs/
COPY --from=0 /cp_bin /bin

COPY peers.json /lachesis_data_dir/
COPY nodes /nodes

# /cp_bin/upx -d /cp_bin/lachesis /cp_bin/crappy_sh /cp_bin/copy /cp_bin/list ;
ENTRYPOINT ["/bin/crappy_sh", "-v", "-e", "-c", "/bin/env ; /bin/list /cp_bin ; /bin/copy /nodes/$node_num/priv_key.pem /lachesis_data_dir/priv_key.pem ; /bin/list /lachesis_data_dir ; /bin/lachesis run /bin/lachesis run --datadir /lachesis_data_dir --store /lachesis_data_dir/badger_db --listen=$node_addr:12000 --heartbeat=100s"]
