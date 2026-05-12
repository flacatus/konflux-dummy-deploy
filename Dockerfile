FROM registry.access.redhat.com/ubi9/go-toolset:1.23 AS build
WORKDIR /opt/app-root/src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /opt/app-root/app main.go

FROM registry.access.redhat.com/ubi9-micro:latest
COPY --from=build /opt/app-root/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
