# 파드 

> 파드라는 단위로 컨테이너를 묶어서 관리하므로 보통 컨테이너 하나가 아닌 여러개의 파드로 구성됨. 

- 파드 생명주기
    - 
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: kubernetes-simple-pod # 파드의 이름 
  labels:
    app: kubernetes-simple-pod # 오브젝트를 식별하는 레이블 
spec:
  initContainers:
  - name: init-myservice 
    image: arisu1000/simple-container-app:latest
    command: ['sh', '-c', 'sleep 2; echo helloworld01;']
  - name: init-mydb
    image: arisu1000/simple-container-app:latest
    command: ['sh', '-c', 'sleep 2; echo helloworld02;']
  containers:
  - name: kubernetes-simple-pod # 컨테이너 이름 
    image: arisu1000/simple-container-app:latest
    resources:
      requests:
        cpu: 0.1
        memory: 200M
      limits:
        cpu: 0.5
        memory: 1G
    ports:
    - containerPort: 8080 # 컨테이너 포트 
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
    env:
    - name: TESTENV01
      value: "testvalue01" # 첫 번째 환경 변수
    - name: HOSTNAME
      valueFrom:
        fieldRef:
          fieldPath: spec.nodeName # 두 번째 환경 변수
    - name: POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name  # 세 번째 환경 변수
    - name: POD_IP
      valueFrom:
        fieldRef:
          fieldPath: status.podIP # 네 번째 환경 변수
    - name: CPU_REQUEST
      valueFrom:
        resourceFieldRef:
          containerName: kubernetes-simple-pod
          resource: requests.cpu # 다섯 번째 환경 변수
    - name: CPU_LIMIT
      valueFrom:
        resourceFieldRef:
          containerName: kubernetes-simple-pod
          resource: limits.cpu # 여섯 번째 환경 변수

```
