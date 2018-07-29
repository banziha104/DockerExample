# 레지스트리

> 도커 이미지를 저장하고 공유할 수 있는 서버

- 오픈소스, Apache 라이센스
- v1과 v2가 호환되지 않는다.
- 클라우드 : DockerHub
- 인트라넷 : DTR

- registry 컨테이너 실행하기
- docker run -d -p 5000:5000 --name myregistry registry:2
- docker tag 8b89e48b5f15 localhost:5000/docker-whale:latest : 태그달기
- docker push localhost:5000/docker-whale:latest : 레지스트리에 등록
- docker pull localhost:5000/docker-whale:latest : 이미지 풀링
- docker run localhost:5000/docker-whale:latest : 이미지 실행

- Storage Drivers : 클라우드 저장소 (s3)
