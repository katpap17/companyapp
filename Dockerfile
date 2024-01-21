FROM golang:1.21 AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /company-app

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /company-app /company-app
EXPOSE 8000

ENTRYPOINT ["/company-app"]