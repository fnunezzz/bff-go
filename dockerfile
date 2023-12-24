FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod .
RUN go mod download && go mod verify

COPY . .

RUN make build-docker

FROM scratch AS deploy-stage

COPY --from=build-stage /app/bin/app /app

CMD ["/app"]