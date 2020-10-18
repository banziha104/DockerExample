# 디플로이먼트 

- 에플리케이션을 다운 타임 없이 업데이트 가능하도록 도와주는 리소스
- 레플리카셋과 레플리케이션커트롤러 상위에 배포되는 리소스
- 스케일링하는법도 레플리카셋과 동일함
- 업데이트 방법론 
    - 모든 포드를 업데이트하는 방법
    - 롤링 업데이트 
    
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy-jenkins
  labels:
    app: jenkins-test
spec:
  replicas: 3
  selector:
    matchLabels:
      app: jenkins-test
  template:
    metadata:
      labels:
        app: jenkins-test
    spec:
      containers:
        - name: jenkins
          image: jenkins
          ports:
            - containerPort: 8080
```

<br>

# 롤링 업데이트와 롤백 

- 새로운 포드를 실행시키고 작업이 완료되면 오래된 포드를 삭제
- 새 버전을 실행하는 동안 구 버전 포드와 연결
- 서비스의 레이블셀렉터를 수정하여 간단하게 수정가능
- 하위 호환성 필요
-  --record=true 옵션을 주어야 확인가능
- 옵션
    - Rolling Update(기본값) 
    - Recreate :이전 버전 다삭제
- 세부 전략
    - maxSurge
        - 기본값 25% 개수로도 설정이 가능
        - 최대로 추가 배포를 허용할 개수 설정
        - 4개인 경우 25%이면 1개가 설정
    - maxUnavailable
        - 기본값 25%, 개수로도 설정 가능
        - 동작하지 않는 포드의 개수 설정
        - 4개인 경우 25%이면 1개가 설정  
- 업데이트를 실패하는 경우
    - 부족한 할당량
    - 레디니스 프로브 실패
    - 이미지 가져오기 오류
    - 권한 부족
    - 제한 범위
    - 응용 프로그램 런타임 구성 오류
    
```shell script

kubectl rollout pause deployment http-go # 일시정지 
kubectl rollout undo deployment http-go # 일시정지 중 취소
kubectl rollout resume deployment http-go # 업데이트 재시작


kubectl rollout status deploy http-go #디플로이 확인 
kubectl rollout history deploy http-go # 디플로이 이력
kubectl patch deploy http-go -p '{"spec":{"minReadySeconds":10}}' #디플로이 변경, 최소 10초간 업데이트 유지
kubectl expose deploy http-go # 서비스 추가 
kubectl get svc # 서비스 확인

kubectl set image deploy http-go http-go=dldudwnsdl/http-go:v2 --record=true # 롤링 업데이트
kubectl edit deploy http-go --record=true # 롤링 업데이트, vim에서 수정 
kubectl rollout undo deploy http-go # 롤백
kubectl rollout undo deploy http-go  --to-revision=1 # 특정 버전으로 돌아감 
``` 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: http-go
  labels:
    app: http-go
spec:
  replicas: 3
  selector:
    matchLabels:
      app: http-go
  template:
    metadata:
      labels:
        app: http-go
    spec:
      containers:
        - name: http-go
          image: dldudwnsdl/http-go:v1
          ports:
            - containerPort: 8080
```