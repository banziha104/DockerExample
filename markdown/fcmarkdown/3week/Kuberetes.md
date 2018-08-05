# 쿠버네티스

> 컨테이너 어플리케이션 배포, 스케일링 , 관리를 자동화 해주는 오픈소스 시스템

- 상태관리 : 노드에 대한 상태를 유지시켜줌
- 스케쥴링 : 조건에 맞는 노드를 찾아서 컨테이너를 배치해준다.
- 클러스터링 : 가상 네트워크를 통해 한대의 서버처럼 묶어준다.
- 서비스 디스커버리
- 리소스 모니터링

### 아키텍쳐 

- Master : 클러스터를 관리, 1대 이상 필요하며 실제로 컨테이너를 배치하지는 않음
    - API 서버 : 클라이언트와 통신하는 HTTPS RestAPI 서버
    - 컨트롤러 매니저 : DaemonSet Controller , Replication Controller 등의 컨트롤러를 관리하고 API 서버와 통신한다.
    - 스케쥴러 : 노드에 리소스 배치한다.
    - etcd : 분산형 key/value 저장소, 설정 정보 저장
    
- Node : 실제로 컨테이너가 실행되는 머신으로 마스터 API 서버와 통신하면서 노드에 컨테이너를 관리한다
    - kubelet : API 서버와 통신하며 컨테이너 상태를 관리
    - 프록시 서버 : 네트워크 프록시, 로드밸런서
    - cAdvisior : 리소스 모니터링
    
- Pod 
    - 배포단위
    - 하나의 Pod은 여러개의 컨테이너를 가짐
    - pod 내에 컨테이너끼리는 로컬콜
    - 고유 IP를 가짐
    - CPU,메모리,볼륨, IP 주소 및 기타 자원을 할당 (pod 내의 모든 컨테이너가 자원을 공유)
    - 컨트롤러로 관리 
    
- Controller 
    - ReplicaSet
        - 팟을 관리함
        - Replication Controller를 대체중
        
    - Deployment (* 최근엔 이거씀)
        - 배포를 관리한다
        - ReplicaSet을 통해 배포한다
    
    - Job : 한번만 실행되고 종료되는 작업을 담당한다.
    - CronJob 
        - 주기적으로 실행되는 작업을 담당한다.
        - Job을 관리한다
        
- Service 
    - pods은 늘어나거나 줄어들 수 있어, pods으로 가는 트래픽을 로드밸런싱해주는 역할
    - Private IP(또는 Public IP)가 Service에 할당
    - Service에 할당된 IP+port를 통해 접근하면, 뒤에 있는 pods에 round-robin 등의 접근 방식으로 접근
    - sticky session 지원
    - ClusterIP : 팟이 유동적이기 때문에 클러스터내 고유한 IP를 만들어 팟간 내부 통신할때 사용한다
    - NodePort : 외부에 접속포트를 오픈할때
    - Loadbalance : AWS ELB 외부 로드밸런서와 연결하는 IP를 만든다
    - Ingress : 모든 노드에 대해서 80/443 포트를 오픈한다

### 설치

- kuberctl 설치

```bash
$ curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl
$ chmod +x ./kubectl
$ sudo mv ./kubectl /usr/local/bin/kubectl

# 아래꺼 안되면 위에꺼 zsh는 이렇게
$ curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl"

```

- 미니쿠베설치

```bash
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

```

### 예제 

1. minikube 시작

```bash
minikube start --vm-driver=virtualbox
```

2. 클러스터 확인

```bash
kubectl cluster-info
```

3. minikube 시작이 안되는 경우 임시파일을 지운다.

```bash
rm -rf ~/.minikube/
```

4. hello-node 이미지 빌드

```bash
eval $(minikube docker-env)
docker build -t hello-node:v1 .

```

5.  hello-node 배포 

```bash
kubectl run hello-node --image=hello-node:v1 --port=8080
```

6. 배포확인

```bash
kubectl get deployments
```

7. pods, events 확인 

```bash
kubectl get pods
```

8. 설정확인

```bash
kubectl config view 
```

9. 서비스 생성

```bash
kubectl expose deployment hello-node --type=LoadBalancer
``` 

10. 서비스 확인

```bash
kubectl get service
```

11. hello-node 테스트

```bash
minikube service hello-node
```

12. 롤링업데이트

```bash
kubectl set image deployment/hello-node hello-node=hello-node:v2
```

13. 대시보드 확인

```bash
kubectl describe svc kubernetes-dashboard -n kube-system
```

14. kube master ip 확인

```bash
minikube status
```

15. Docker환경 테스트

```bash
minikube docker-env 
```

# YAML 이용

- nginx-pod.yaml

```yaml

apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.7.9
    ports:
    - containerPort: 80
    
```

- 명령어 실행

```yaml
kubectl create -f nginx-pod.yaml 
```

- redis-pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  containers:
  - name: redis
    image: redis:5.0-rc4
    volumeMounts:
    - name: redis-storage
      mountPath: /data/redis
  volumes:
  - name: redis-storage
    emptyDir: {}
```

- git-monitor.yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: www
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - mountPath: /srv/www
      name: www-data
      readOnly: true
  - name: git-monitor
    image: kubernetes/git-monitor
    env:
    - name: GIT_REPO
      value: https://github.com/k16wire/docker-whale
    volumeMounts:
    - mountPath: /data
      name: www-data
  volumes:
  - name: www-data
    emptyDir: {}
```


- nginx-deploymenet.yaml

```yaml
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.8.0
        ports:
        - containerPort: 80
```

```bash
kubectl create -f nginx-deployment.yaml
```

### 기타


