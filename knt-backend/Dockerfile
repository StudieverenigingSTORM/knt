# We can't use golang alpine (CGO disabled), we need glibc for sqlite3
FROM golang:1.20 AS build
WORKDIR /app
COPY ./go.mod ./go.sum ./main.go ./
COPY ./internal ./internal
# We cannot build w/ CGO disabled or static bin, as we require glibc for sqlite3
# -s & -w: strips debug info 
RUN go build -ldflags "-s -w" -o /knt .

# Distroless: tiny images, only runtime deps, by Google. See https://github.com/GoogleContainerTools/distroless
# CGO forces us to use base (incl. glibc) instead of static
FROM gcr.io/distroless/base-debian11 AS release
WORKDIR /
COPY --from=build /knt /knt
COPY ./config ./config
EXPOSE 5000
ENTRYPOINT ["/knt"]
