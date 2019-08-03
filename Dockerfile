FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /go/src/main
ENV GO111MODULE=on
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /app /app
# Run the hello binary.
ENTRYPOINT ["/app"]