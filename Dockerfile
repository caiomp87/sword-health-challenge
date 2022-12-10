FROM debian:buster

RUN apt-get update \
    && apt-get install -y ca-certificates --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && update-ca-certificates

RUN addgroup kwan \
    && adduser --disabled-password --gecos "" kwan --uid 1000 --ingroup kwan \
    && chown kwan:root /etc/ssl/private

USER kwan

COPY --chown=kwan:kwan app /usr/local/bin/

RUN chmod +x /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]
