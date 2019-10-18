############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/slack-random-emo
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN go build -o /go/bin/slack-random-emo

############################
# STEP 2 build a small image
############################
FROM alpine:3.10

# Copy our static executable.
COPY --from=builder /go/bin/slack-random-emo /go/bin/slack-random-emo

# Run the hello binary.
ENTRYPOINT ["/go/bin/slack-random-emo"]
