# 한 포드에 멀티 컨테이너를 사용 

- 이 포드는 볼륨, IPC, 네트워크 인터페이스를 공유, 지역성을 보장하고, 여러 개의 응용프로그램이 결합된 형태로 하나의 포드를 구성할 수 있음.
- 주로 리소스 모니터링용도로 사용으로 많이 사용됨.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-redis-pod
spec:
  containers:
    - name: nginx
      image: nginx
      ports:
        - containerPort: 80
    - name: redis
      image: redis
```

<br>

# initContainer

- 포드 컨테이너 실행 전에 초기화 역할을 하는 컨테이너
- 완전히 초기화가 진행된 다음에야 주 컨테이너를 실행
- init 컨테이너가 실행하면, 성공할때가지 포드를 반복해서 재시작
- restartPolicy에 Never를 하면 재시작하지 않음.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox:1.28
    command: ['sh', '-c', 'echo The app is running! && sleep 3600']
  initContainers:
  - name: init-myservice
    image: busybox:1.28
    command: ['sh', '-c', "until nslookup myservice.$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace).svc.cluster.local; do echo waiting for myservice; sleep 2; done"]
  - name: init-mydb
    image: busybox:1.28
    command: ['sh', '-c', "until nslookup mydb.$(cat /var/run/secrets/kubernetes.io/serviceaccount/namespace).svc.cluster.local; do echo waiting for mydb; sleep 2; done"]
```