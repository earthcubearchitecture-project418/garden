# Start from scratch image and add in a precompiled binary
# CGO_ENABLED=0 env go build .
# docker build -t earthcube/p0servicelp:latest -t earthcube/p0servicelp:0.0.1 .
# docker run -d -p 7777:7777 -p 6789:6789  earthcube/p0servicelp:latest
FROM scratch

# Add in the static elements (could also mount these from local filesystem) later as the indexes grow
ADD serviceLP /
ADD ./static /static

# Add our binary
CMD ["/serviceLP"]

# Document that the service listens on this port
EXPOSE 6789 7777
