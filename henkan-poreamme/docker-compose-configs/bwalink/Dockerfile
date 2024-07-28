FROM ruby:3-alpine

# Fix for dev tools from https://renehernandez.io/snippets/install-development-tools-in-alpine/
# This also adds the build-base, uses it and then wipes out the tools to keep the image small
# Installing socat to link the serial over IP device
RUN apk update && apk add --no-cache build-base && \
    bundle config set deployment 'true' && \
    gem install balboa_worldwide_app && \
    apk add --no-cache socat tzdata && \
    apk del build-base && \
    rm -rf /var/cache/apk/*

ENTRYPOINT /usr/local/bundle/bin/bwa_mqtt_bridge ${MQTT_URI} tcp://${BRIDGE_IP}:${BRIDGE_PORT}/

