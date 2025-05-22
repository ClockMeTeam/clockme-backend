FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*



COPY ./ftfclockify ./bin/ftfclockify
COPY ./.env ./.env

CMD ["/bin/ftfclockify"]