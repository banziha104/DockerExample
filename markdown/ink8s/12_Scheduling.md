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

<br>

# 시스템 리소스 요구사항과 제한 설정 

- CPU 및 메모리는 각각 자원 유형을 지니며 자원 유형에는 기본 단위를 사용 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      run: nginx
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        run: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
        ports:
        - containerPort: 80
        resources: 
          requests:
            memory: "200Mi"
            cpu: "1m"
          limits:
            memory: "400Mi"
            cpu: "2m"
status: {}
```

<br>

# 리미트 레인지

- 네임 스페이스에서 포드나 컨테이너당 최소 및 최대 컴퓨팅 리소스 사용량 제한
- 네임 스페이스에서 PersistentVolumeClaim 당 최소 및 최대 스토리지 사용량 제한
- 네임 스페이스에서 리소스에 대한 요청과 제한 사이의 비율 적용
- 네임 스페이스에서 컴퓨팅 리소스에 대한 디폴트 requests/limit를 설정하고 럼타임 중인 컨테이너에 자동으로 입력

```shell
--enable-admission-plugins=LimitRange
```

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: mem-limit-range
spec:
  limits:
    - default:
        memory: "512Mi"
      defaultRequest:
        memory: "256Mi"
      max:
        cpu: "800m"
        memory: "1Gi"
      min:
        cpu: "100m"
        memory: "99Mi"
      type: Container

```

<br>

# ResourceQuata

> 네임스페이스별 총량 리소스 제한

```yaml
apiVersion: v1
  kind: ResourceQuota
  metadata:
    name: pods-low
  spec:
    hard:
      cpu: "5"
      memory: 10Gi
      pods: "10"
    scopeSelector:
      matchExpressions:
        - operator: In
          scopeName: PriorityClass
          values: [ "low" ]
```

<br>

# 데몬셋 

- 레플리케이션컨트롤러와 레플리카셋은 무작위 노드에 포드를 생성
- 데몬셋은 각 하나의 노드에 하나의 포드만을 구성
- kube-proxy가 데몬셋으로 만든 쿠버네티스에서 기본적으로 활동중인 포드

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: http-go
spec:
  selector:
    matchLabels:
      name: http-go
  template:
    metadata:
      labels:
        app: http-go
    spec:
      tolerations:
      - key: node-role.kubernetes.io/master
        effect: NoSchedule
      containers:
      - name: http-go
        image: gasbugs/http-go
```

<br>

# 스태틱 포드

> 마스터 서버에 존재하고, kubelet이 직접 실행하는 포드 

- 각각의 노드에서 kubelet에 의해 실행
- 포드를 삭제할때, apiserver를 통해서 실행되지 않은 스태틱 포드는 삭제 불가
- 노드에 필요에 의해 사용하고자 하는 포드는 스태틱 포드로 세팅

```shell

# master server
sudo -i
cd /etc/kubernetes/manifests/
vim <static-pod-name>.yaml

```

<br>

# 수동 스케줄링 

- 수동으로 원하는 노드에다가 배치하는 방식
- 특수한 환경의 경우 특정 노드에서 실행되도록 선호하도록 포드를 제한
- 일반적으로 스케줄러는 자동으로 합리적인 배치를 수행하므로 이러한 제한은 필요없음
- 특수 케이스
    - SSD가 있는 노드에서 포가 실행되기 위한 경우
    - 블록체인이나 딥러닝 시스템을 위해 GPU가 서비스가 필요한 경우
    - 서비스의 성능을 극대화하기 위해 하나의 노드에 필요한 포드를 모두 배치

- yaml 파일 활용

```yaml
apiVersion: apps/v1
kind: Pod
metadata:
  name: http-go
spec:
  containers:
    - name: http-go
      image: busybox
  nodeName: work1
```

- 노드 셀렉터 활용

```shell
kubectl label node <node_name> gpu=true
```

```yaml
apiVersion: apps/v1
kind: Pod
metadata:
  name: http-go
spec:
  containers:
    - name: http-go
      image: busybox
  nodeSelector:
    gpu: "true"
```

<br>

# 멀티플 스케줄러