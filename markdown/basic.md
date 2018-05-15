# 도커기본 사용법

- 검색
    - sudo docker search nginx : search 명령으로 특 이미지 검색

- 이미지 받기
    - sudo docker pull ubuntu:latest : pull 명령으로 이미지 받기
    - docker images : 받은 이미지 목록 출력
    - docker run -i -t --name hello ubuntu /bin/bash : run 명령어로 컨테이너 생성
    - docker run <옵션> <이미지 이름> <실행할 파일>
    - docker ps -a : 컨테이너 목록 확인
        - a : 정지된 컨테이너 까지 모두 출력
    - docker start <컨테이너 이름>  : 컨테이너 실행
    - docker restart <컨테이너 이름> : 컨테이너 재시작
    - docker attach <컨테이너 이름> : 컨테이너에 접속
    - docker exec <컨테이너 이름>  <명령> <매개변수> : 외부에서 컨테이너 안의 명령 실행
    - docker stop <컨테이너> : 컨테이너 정지
    - docker rm <켄터이너 이름> : 컨테이너 삭제 
    - docker rmi <이미지 이름>:<태그> : 이미지 삭제