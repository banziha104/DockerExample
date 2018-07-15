# 아마존 리눅스 설치 

- amazon-linux-extras install docker : Amazon 리눅스에 도커설치

- service docker start : 도커 데몬 실행

- docker pull nginx:latest : Nginx를 받아옮

- docker run -d -p 8080:80 --name Nginx nginx:latest : 호스트의 8080 번 포트를 컨테이너의 80 포트와 연결하고 Nginx라는 이름으로 구동

- docker ps -a : 도커 실행 확인

