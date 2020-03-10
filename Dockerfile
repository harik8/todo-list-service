FROM golang:1.13 as build

LABEL HARI KARTHIGASU <hariprasad.karthigasu@gmail.com>

WORKDIR /app

COPY go.mod go.sum ./ \
     . ./

RUN go build .

FROM gcr.io/distroless/base

WORKDIR /app

EXPOSE 8080

COPY --from=build /app/todo-list-service .

CMD ["todo-list-service"]