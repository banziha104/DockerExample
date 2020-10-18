# ReplicaController

- ReplicaContronller : ReplicatSet의 구버전
    - 레플리케이션 : 데이터 저장과 백업하는 방법과 관련이 있는 데이터를 호스트 컴퓨터에서 다른 컴퓨터로 복사하는 것 
    - 포드가 항상 실행되도록 유지하는 쿠버네티스 리소스
    - 노드가 클러스터에서 사라지는 경우 해당 포드를 감지하고 대체 포드 생성
    - 실행 중인 포드의 목록을 지속적으로 모니터링하고 '유형'의 실제 포드 수가 원하는 수와 항상 일치하는지 확인 
 - 요소
    - 포드 범위를 결정하는 레이블 셀렉터
    - 실행해야 하는 포드의 수를 결정하는 본제본 수
    - 새로운 포드의 모양을 설명하는 포드 템플릿
- 장점
    - 포드가 없는 경우 새포드를 항상 실행
    - 노드에 장애 발생 시 다른 노드에 복제본 생성
    - 수동,자동 스케일링
    
```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx
spec:
  replicas: 3
  selector:
    app: nginx
  template:
    metadata:
      name: nginx
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
```

- 스케일 

```shell script
kubectl scale rc http-go --replicas=5 
kubectl edit rc http-go # vim으로 수정
kubectl apply -f http-go.yaml # YAML 파일로 수정해서 올림
```

<br>

# ReplicaSet

- 레플리카셋은 차세대 레플리케이션 컨트롤러로 대체 가능
- 1.8 버전부터 디플로이먼트,데몬셋,레프리카셋,스트레이프 풀셋이 베타, 19에서는 정식으로 업데이트
- Replication Controller 
    - 거의 동일하게 동작
    - 더 풍부한 표현식 포드 셀렉터 사용 가능 
    - 레플리케이션 컨트롤러 : 특정 레이블을 퐇마하는 포드가 일치하는지 확인
    - 레플리카셋: 특정 레이블이 없거나 해당 값과 관계없이 특정 레이블 키를 포함하는 포드를 매치하는지 확인
    - rc 대신 rs 만 사용하면 커맨드에서 동일하게 사용가능 
    
```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: rs-nginx
spec:
  # modify replicas according to your case
  replicas: 3
  selector:
    matchLabels:
      app: rs-nginx
  template:
    metadata:
      labels:
        app: rs-nginx
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - containerPort: 80
```