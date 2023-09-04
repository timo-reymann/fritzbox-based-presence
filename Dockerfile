FROM busybox AS bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/fritzbox-based-presence_linux-${architecture} /bin/fritzbox-based-presence && \
    chmod +x /bin/fritzbox-based-presence && \
    chown 1000:1000 /bin/fritzbox-based-presence

FROM gcr.io/distroless/static-debian11:nonroot
LABEL org.opencontainers.image.title="fritzbox-based-presence"
LABEL org.opencontainers.image.description="Show who is home based on devices connected to FritzBox that are currently online."
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GNU GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/fritzbox-based-presence"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/fritzbox-based-presence"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/fritzbox-based-presence.git"
COPY --from=bin /bin/fritzbox-based-presence /bin/fritzbox-based-presence
ENTRYPOINT ["/bin/fritzbox-based-presence"]
