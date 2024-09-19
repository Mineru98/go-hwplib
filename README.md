# Nginx 여러가지 모듈 실험

## nginx에서 log

기본적으로 docker에서 volumn 연결을 하지 않으면 log는 /dev/stdout 으로 쌓이고 
error log는 /dev/stderr에 쌓이게 됩니다.
장치 폴더(/dev)에에 stdout은 버퍼 역할을 하는 파일이며 프로세스가 종료되면 사라집니다.

만약 docker에 volumn 연결을 했다면, 실제 모든 access.log와 error.log가 해당 인스턴스에 모두 쌓이게 됩니다.
만약 이렇게 해야하는 경우라면 logrotate를 사용하여 주기적으로 로그를 쌓고 용량 제한을 두어 관리를 해야 합니다.
