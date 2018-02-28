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
    

---

# Bash 익히기

- \> : 출력 다이렉션, 덮어씀
- \< : 입력 다이렉션
- \>> : 명령 실행의 표춘 출력을 파일에 추가, 뒷부분에 추가함 
- 2> : 표준 에러를 파일로 저장
- 2>> : 표준 에러를 파일에 추가
- &> : 표준 출력과 표준 에러를 모두 파일로 저장함

- 1>&2 : 표준 출력을 표준 에러로 보냄.
- | : 명령 실행의 표준 출력을 다른 명령의 표준 입력으로 보냄
- $ : 변수사용시에만 씀
- $() : 명령 실행 결과를 변수화함. 명령 실행 결과를 변수에 저장하거나 다른 명령의 매개변수로 넘겨줄 때 사용
- `` : 명령 실행결과를 변수화함 $()와 같음
- 


```dockerfile

```
    