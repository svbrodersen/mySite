# Build
FROM golang:1.23rc1 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Deploy.
FROM gcr.io/distroless/static-debian11 AS release-stage
WORKDIR /
COPY --chown=nonroot --from=build-stage /entrypoint /entrypoint
COPY --chown=nonroot --from=build-stage /app/static /static
COPY --chown=nonroot --from=build-stage /app/projects.db /projects.db
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/entrypoint"]
