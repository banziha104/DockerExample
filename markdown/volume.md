## 도커 볼륨

> 이미지는 읽기 전용이 되며, 변경사항만 별도로 저장해서 각 컨테이너의 정보를 보존
> <br> 컨테이너를 삭제하면 데이터베이스의 정보도 삭제됌, 볼륨을 ㅘㄹ용하면 데이터를 영속적으로 활용가능

- sudo docker run -i -t --name hello-volume -v /data ubuntu /bin/bash
    - 호스트 볼륨 사용
    - -v <컨테이너 디렉터리>로 볼륨 설정
    - /root/data로 데이터가 모임
    
- sudo docker run -i -t --name hello-volume -v /root/data:/data ubuntu /bin/dash
    - 데이터 볼륨 사용
    - -v <호스트 디렉터리> : <컨테이너 디렉터리> 
    