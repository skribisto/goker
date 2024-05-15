##################################################################################
FROM --platform=$BUILDPLATFORM golang:bookworm AS goker-builder

# Install needed binaries
RUN apt-get update && apt-get install -y build-essential crossbuild-essential-armhf crossbuild-essential-armel crossbuild-essential-arm64 crossbuild-essential-i386

# Prepare the source location
RUN mkdir -p /go/src/github.com/skribisto/goker
WORKDIR /go/src/github.com/skribisto/goker

RUN mkdir release

# Add the source code ( see .dockerignore )
COPY . .

RUN go build -o release/goker-v0.1.1

##################################################################################
FROM alpine:3.19 AS goker-image

RUN apk add --no-cache ca-certificates

# Create goker user
ENV USER=goker
ENV UID=1000

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/home/goker" \
    --shell "/bin/false" \
    --uid "${UID}" \
    "${USER}"

COPY --from=goker-builder --chown=1000:1000 /go/src/github.com/skribisto/goker/release /home/goker/

USER goker
WORKDIR /home/goker
CMD ./goker-v0.1.1