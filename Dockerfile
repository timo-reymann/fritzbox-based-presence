# Start by building the application.
FROM golang:1.21 as build

ARG BUILD_TIME
ARG BUILD_VERSION
ARG BUILD_COMMIT_REF
LABEL org.opencontainers.image.created=$BUILD_TIME
LABEL org.opencontainers.image.version=$BUILD_VERSION
LABEL org.opencontainers.image.revision=$BUILD_COMMIT_REF

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/app /
CMD ["/app"]