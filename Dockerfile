# Build the api and ingest binaries, then ship them in a distroless image.
# The web/ UI is embedded into the api binary via go:embed, so no assets are
# copied into the final stage.

FROM golang:1.26 AS build
WORKDIR /src

# Cache module downloads.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api   ./cmd/api  && \
    CGO_ENABLED=0 GOOS=linux go build -o /out/ingest ./cmd/ingest

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/api   /app/api
COPY --from=build /out/ingest /app/ingest
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app/api"]
