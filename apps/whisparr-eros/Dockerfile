# syntax=docker/dockerfile:1

FROM docker.io/library/alpine:3.24
ARG TARGETARCH
ARG VENDOR
ARG VERSION

# krb5 on musl isn't thread-safe and Npgsql's default GSS probe segfaults on it (exit 139).
ENV COMPlus_EnableDiagnostics=0 \
    PGGSSENCMODE=disable

USER root
WORKDIR /app

RUN \
    apk add --no-cache --upgrade \
        bash \
        ca-certificates \
        catatonit \
        coreutils \
        curl \
        icu-libs \
        jq \
        krb5-libs \
        libintl \
        nano \
        sqlite-libs \
        tzdata \
        vips \
    && mkdir -p /app/bin \
    && curl -fsSL -o /tmp/whisparr.tar.gz "https://github.com/Whisparr/Whisparr-Eros/releases/download/${VERSION}/Whisparr.eros-develop.${VERSION#v}.linux-musl-${TARGETARCH/amd64/x64}.tar.gz" \
    && tar xzf /tmp/whisparr.tar.gz -C /app/bin --strip-components=1 \
    && printf "UpdateMethod=docker\nBranch=%s\nPackageVersion=%s\nPackageAuthor=[%s](https://github.com/%s)\n" "eros-develop" "${VERSION#v}" "${VENDOR}" "${VENDOR}" > /app/package_info \
    && chown -R root:root /app && chmod -R 755 /app \
    && rm -rf /tmp/* /app/bin/Whisparr.Update

COPY --chmod=755 entrypoint.sh /entrypoint.sh

USER nobody:nogroup
WORKDIR /config
VOLUME ["/config"]

ENTRYPOINT ["/usr/bin/catatonit", "--", "/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/Whisparr/Whisparr-Eros"
