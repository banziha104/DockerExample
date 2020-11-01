# 어플리케이션 변수 관리

- 쿠버네티스의 환경 변수는 YAML 파일이나 다른 리소스로 전달

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: envar-demo
  labels:
    purpose: demonstrate-envars
spec:
  containers:
  - name: envar-demo-container
    image: gcr.io/google-samples/node-hello:1.0
    env:
    - name: DEMO_GREETING
      value: "Hello from the environment"
    - name: DEMO_FAREWELL
      value: "Such a sweet sorrow"
```
```yaml
env:
  - name: SERVER_ADDRESS
    value: http://server.com
  - name: SERVER_ADDRESS
    valueFrom: 
      configMapKeyRef: configmap-name
  - name: SERVER_ADDRESS
    valueFrom:
      secretKeyRef: secret-name
```

<br>

# ConfigMap

> 환경 변수를 저장하거나 스토리지를 저장하는 리소스 

```shell
echo -n 1234 > test 
kubectl create configmap map-name --from-file=test
```

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: special-config
  namespace: default
data:
  SPECIAL_LEVEL: very
  SPECIAL_TYPE: charm
```


<br>

# Secret

> 컨피그맵은 평문으로, Secret은 인코딩으로 저장됨

```shell
echo -n admin > username
echo -n asdf > password
kubectl create secret generic db-user-pass --from-file=username --from-file=passwordx
```

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
type: Opaque
data:
  username: YWRTaW4=
  password: MWYyZDF1MmU2N2Rm
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: envar-secret
  labels:
    purpose: demonstrate-envars
spec:
  containers:
  - name: envar-demo-container
    image: gcr.io/google-samples/node-hello:1.0
    env:
    - name: user
      valueFrom:
        secretKeyRef:
          name: db-user-p ass
          key: username
    - name: pass
      valueFrom:
        secretKeyRef:
          name: db-user-pass
          key: password
```

<br>

# 초기 명령어 및 아규먼트 전달과 실행 

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox
  labels:
    app: myapp
spec:
  containers:
  env:
    - name: MESSAGE
      value: "hello world"
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo Hello Kubernetes! && sleep 3600']
    args: ["HOSTNAME", "$(MESSAGE)"]
```