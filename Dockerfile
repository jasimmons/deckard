FROM ubuntu:18.04

ADD bin/deckard-linux-amd64 /usr/bin/deckard

ENTRYPOINT ["/usr/bin/deckard", "serve"]
