FROM nginx:1.27.1-alpine

# logrotate 및 cron 설치
RUN apk update && apk add --no-cache logrotate dcron

# logrotate 설정 파일 복사
COPY nginx-logrotate /etc/logrotate.d/nginx

# logrotate를 매일 실행할 cronjob 설정
COPY logrotate-cron.sh /etc/periodic/daily/logrotate

# 실행 권한 설정
RUN chmod +x /etc/periodic/daily/logrotate

# Nginx 로그 디렉토리 생성 및 권한 설정 (필요시)
RUN mkdir -p /var/log/nginx && \
    chown nginx:nginx /var/log/nginx

EXPOSE 80

CMD ["sh", "-c", "crond & nginx -g 'daemon off;'"]