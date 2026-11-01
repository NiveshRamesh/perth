# syntax=docker/dockerfile:1
# Tip: pin base images by digest and let Dependabot bump them.

FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/perth .

# distroless/static: no shell, no package manager, runs as uid 65532 - tiny attack surface
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/perth /perth
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/perth"]
