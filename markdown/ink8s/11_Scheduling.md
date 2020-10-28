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


