FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY mountebank_merge_stubs.go ./

# Initialize Go module
RUN go mod init mountebank_merge_stubs && \
    go build -o mountebank_merge_stubs

FROM node:23-alpine

WORKDIR /app

COPY --from=builder /build/mountebank_merge_stubs ./
COPY stubs ./stubs

RUN npm install -g mountebank@2.9.1

RUN chmod +x ./mountebank_merge_stubs && \
    ./mountebank_merge_stubs

EXPOSE 2525 8090

CMD ["mb", "--configfile", "imposter.json", "--allowInjection"]