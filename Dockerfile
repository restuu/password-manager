################
# BUILD BINARY #
################
#  golang:1.15.4-alpine3.12
FROM golang@sha256:fac295af2c87baca24ec25efebddfdfc48bb0abb63a51df8eab351c20630fa61 as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Fetch dependencies.
RUN go mod download
RUN go mod verify

COPY . .

RUN echo $PWD && ls -lah

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o ./bin/password-manager .

#####################
# MAKE SMALL BINARY #
#####################
FROM alpine:3.11

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /app

# Copy the executable.
COPY --from=builder /app/bin/password-manager .

RUN echo $PWD && ls -lah

CMD ["./password-manager", "server"]