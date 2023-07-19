FROM golang:1.19-alpine3.16 AS server_builder

RUN apk add gcc libc-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . TestTaskHaloLab/.

WORKDIR TestTaskHaloLab
RUN go build -ldflags "-w -s -linkmode external -extldflags -static" -a cmd/app/main.go

FROM scratch
EXPOSE 8080
COPY --from=server_builder /app/TestTaskHaloLab/main .
COPY --from=server_builder /app/TestTaskHaloLab/.env .
CMD ["./main"]