FROM golang:buster as builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 \
	GOOS=linux \
    GOARCH=amd64 \
	go build -o pbbot_app_loginhelper ./cmd/loginhelper/

FROM debian:bullseye as runner
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY --from=builder /app/pbbot_app_loginhelper .
COPY --from=builder /app/cacert.pem /etc/ssl/certs/cacert.pem
VOLUME /app/conf
VOLUME /app/autocert
VOLUME /app/log
EXPOSE 8081

ENTRYPOINT ["./pbbot_app_loginhelper" ,"-c","/app/conf/config.yaml"]
