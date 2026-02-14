FROM busybox AS bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/fritzbox-based-presence_linux-${architecture} /bin/fritzbox-based-presence && \
    chmod +x /bin/fritzbox-based-presence && \
    chown 65532:65532 /bin/fritzbox-based-presence

FROM scratch
LABEL 
LABEL org.opencontainers.image.licenses="gpl-3.0"org.opencontainers.image.title="fritzbox-based-presence"
LABEL org.opencontainers.image.description="Show who is home based on devices connected to FritzBox that are currently online."
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GNU GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/fritzbox-based-presence"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/fritzbox-based-presence"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/fritzbox-based-presence.git"
COPY --from=gcr.io/distroless/static-debian12:nonroot / /
USER nonroot
COPY --from=bin /bin/fritzbox-based-presence /bin/fritzbox-based-presence
ENTRYPOINT ["/bin/fritzbox-based-presence"]

