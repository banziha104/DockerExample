# Docker Security

- 도커 호스트 안전
    - 도커 데몬은 root 권한으로 실행됨
    - 도커 데몬은 신뢰할 수 있는 사용자만 접근가능해야 한다
    - REST API는 TCP 소켓대신 UNIX 소켓을 사용한다
    - 도커 호스트 서버에 어드민 관리 도구를 실행하지 말것
    - --rm으로 띄우고 나올때 삭제해버리기
    
- 컨테이너가 해킹당할 경우
    - 컨테이너에서 실행되는 프로세스는 다른 컨테이너의 프로세스에 영향을 줄 수 없다.
    - 컨테이너는 자신만의 네트워크 스택을 갖는다
    - 브릿지 
    
- 컨테이너가 폭주하면
    - 컨테이너는 메모리, CPU , I/O 자원을 공유한다.
    - 컨테이너는 호스트 서버 모든 자원을 소진할 수 없다.
    - Kernal panic과 DoS 방지
    
- 도커 이미지 안전
    - 도커 이미지는 안전한 위치에 저장되어야 하며 체크 섬을 가져야한다.
    - 배포전에 보안패치를 빌드해야 한다.
    - --privileged 사용하지 않는다 
    
- Docker Security Scanning : 레지스트리에 애드온되어 이미지 보안 취약점을 스캔해서 알려줌 
- Docker Bench : 도커 호스트와 컨테이너에 대한 보안 취약점을 체크해주는 스크립트

```bash
git clone https://github.com/docker/docker-bench-security
cd docker-bench-security
docker-compose run --rm docker-bench-security # 실행
``` 