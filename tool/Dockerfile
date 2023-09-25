#- Lint -#
FROM golangci/golangci-lint:v1.53.3 as lint
WORKDIR /app
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
     golangci-lint run .


#- Build -#
FROM golang as build
WORKDIR /app
RUN --mount=target=. --mount=target=/root/.cache,type=cache \
    --mount=target=/go/pkg/mod,type=cache \
    go build -o /demo .
ENTRYPOINT ["./demo"]


#- Local build binary -#
FROM scratch as binary
COPY --from=build /demo /demo
