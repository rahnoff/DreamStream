# syntax=docker/dockerfile:1
FROM golang:1.25-alpine@sha256:301db1784c200c7ec221a84672d0bd3ffdcf6c873bc9806482bbe730910924fb AS build
WORKDIR /opt/dream-stream
COPY ./go.mod ./
RUN go mod download
COPY ./*.go ./
RUN go build -o ./courses \
    && gropadd -r courses \
    && useradd -c 'Courses user' -d /opt/dream-stream -g courses -r -s /sbin/nologin courses

FROM scratch
WORKDIR /opt/dream-stream
COPY --from=build /etc/group /etc/group \
    --from=build /etc/passwd /etc/passwd \
    --from=build /opt/dream-stream/courses ./
ENTRYPOINT ["/opt/dream-stream/courses"]
