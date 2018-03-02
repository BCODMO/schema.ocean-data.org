# Start from scratch image and add in a precompiled binary
# docker build --tag="BCODMO/schemaodo:latest" --tag="BCODMO/schemaodo:" .
# docker run -d -p 9900:9900  BCODMO/schemaodo:latest
# docker save earthcube/p418vocab:latest | bzip2 | ssh root@geodex.org 'bunzip2 | docker load'
FROM scratch

# Add in the static elements (could also mount these from local filesystem)
# later as the indexes grow
ADD server /
ADD ./html/ ./html

# Add our binary
CMD ["/server"]

# Document that the service listens on this port
EXPOSE 9900
