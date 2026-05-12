FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /app main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
