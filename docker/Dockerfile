FROM golang:1.26 AS builder
ARG SVC
# TARGETARCH / TARGETVARIANT are set automatically by docker buildx.
# GOARCH / GOARM can still be passed explicitly via --build-arg to keep
# backwards-compatibility with the `make dockers` Makefile targets.
ARG TARGETARCH
ARG TARGETVARIANT
ARG GOARCH=${TARGETARCH}
ARG GOARM
ARG VERSION
ARG COMMIT
ARG TIME

WORKDIR /go/src/github.com/absmach/propeller
COPY . .
RUN apt-get update \
    && apt-get install -y --no-install-recommends make \
    && _goarm="${GOARM:-${TARGETVARIANT#v}}" \
    && make $SVC GOARCH=${GOARCH} ${_goarm:+GOARM=${_goarm}} \
    && mv build/$SVC /exe

# distroless/cc-debian12 has no riscv64 manifest; fall back to debian:sid-slim
# for that platform, which provides glibc and libgcc required by CGO binaries.
FROM gcr.io/distroless/cc-debian12 AS final-amd64
FROM gcr.io/distroless/cc-debian12 AS final-arm64
FROM gcr.io/distroless/cc-debian12 AS final-arm
FROM gcr.io/distroless/cc-debian12 AS final-ppc64le
FROM gcr.io/distroless/cc-debian12 AS final-s390x
FROM debian:sid-slim AS final-riscv64
RUN apt-get update \
    && apt-get install -y --no-install-recommends libgcc-s1 \
    && rm -rf /var/lib/apt/lists/*

FROM final-${TARGETARCH}
LABEL org.opencontainers.image.source=https://github.com/absmach/propeller
LABEL org.opencontainers.image.licenses=Apache-2.0

COPY --from=builder /exe /
ENTRYPOINT ["/exe"]
