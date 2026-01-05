# syntax=docker/dockerfile:1
FROM golang:1.25-alpine@sha256:301db1784c200c7ec221a84672d0bd3ffdcf6c873bc9806482bbe730910924fb AS build
WORKDIR /opt/dream-stream
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./*.go ./
RUN go build -o ./courses \
    && addgroup -S courses \
    && adduser -g 'Courses user' -h /opt/dream-stream -s /sbin/nologin -D -G courses -H -S courses

FROM scratch
WORKDIR /opt/dream-stream
COPY --from=build /etc/group /etc/group
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /opt/dream-stream/courses ./
USER courses
ENTRYPOINT ["/opt/dream-stream/courses"]
