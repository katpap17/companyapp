FROM golang:1.21 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /company-app

FROM alpine:latest AS build-release-stage
WORKDIR /
COPY --from=build-stage /company-app /company-app
COPY --from=build-stage /app/config.json /config.json
EXPOSE 8000

ENTRYPOINT ["/company-app"]