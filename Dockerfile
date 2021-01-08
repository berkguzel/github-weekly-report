FROM golang:1.14.13-alpine AS builder
WORKDIR /src/github-weekly-report
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 go build -o  /app ./cmd

FROM alpine
COPY --from=builder /app ./github-weekly-report
ENTRYPOINT [ "./github-weekly-report" ]