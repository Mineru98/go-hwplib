FROM nginx:1.27.1

# logrotate 설치
RUN apt-get update && apt-get install -y logrotate

# logrotate 설정 파일 복사
COPY nginx-logrotate /etc/logrotate.d/nginx

# logrotate를 매일 실행할 cronjob 설정
COPY logrotate-cron.sh /etc/cron.daily/logrotate

# Nginx 로그 디렉토리 생성 및 권한 설정 (필요시)
RUN mkdir -p /var/log/nginx && \
    chown www-data:www-data /var/log/nginx

EXPOSE 80

CMD service cron start && nginx -g 'daemon off;'
