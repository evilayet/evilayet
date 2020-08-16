FROM golang:alpine AS builder

ENV GOOS linux
ENV CG0_ENABLED 0
ENV HOME /app

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/tzone -v main.go

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=builder /app/bin/tzone /app/bin/tzone
COPY --from=builder /app/province.json /app/bin/province.json
COPY --from=builder /app/county.json /app/bin/county.json
COPY --from=builder /app/village.json /app/bin/village.json
COPY --from=builder /app/towns.json /app/bin/towns.json

WORKDIR /app/bin

EXPOSE 12071

ENTRYPOINT ["./tzone"]

#CMD ["sh", "-c", "tail -f /dev/null"]
