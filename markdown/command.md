# 도커 살펴보기 및 히스토리

- docker history hello:0.1 : 히스토리 찾기
- docker cp hello:/etc/nginx/nginx.conf ./ : 호스트 경로로 복사
- docker commit -a "<foo@bar.com>" -m "add hello.txt" hello-nginx hello0.2 : 변경사항을 이미지로 만듬
- docker diff hello-nginx : 변경된 파일 확인
- docker inspect hello-nginx : 세부정보 확인

