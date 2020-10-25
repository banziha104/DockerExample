# 쿠버네티스 

 - Master Node  (컨트롤 플레인)
    - etcd : 데이터베이스 역할 
    - kube-apiserver : 중앙에서 관리하는 역활(사용자 인증, 통신, 권한 등)
    - kube-scheduler : 스케쥴링, 배치 전략
    - controller-manager : 관련된 리소스(YAML)을 컨트롤하는 역할
- Worker Node (노드)
    - kubelet : 워커 노드에서 kube-apiserver로 부터 명령을 하달받고 행동 (주로 컨테이너 런타임을 조절)
    - kube-proxy : 통신과 관련된 역할을 함. 
    
- 쿠버네티스를 관리하는 명령
    - kubeadm : 클러스터를 부트스트랩하는 명령
    - kubelet : 클러스트의 모든 시스템에서 실행되는 구성 요소로, 창 및 컨테이너 시작과 같은 작업을 수행
    - kubectl : 커맨드 라인, 클라이언트 프로그램 

<br>

## 설치 

- 스왑을 비활성화하는 이유
    - k8s 1.8 이후 노드에서 스왑을 비활성화해야됨
    - 스왑 발생시 속도가 느려지는 이슈
    - 스케줄러가 포드를 머신에 보내면 스왑을 사용하지 않는 것이 필요.
    

```shell script

# 도커 설치 
sudo apt install docker.io

# 스크립트 파일을 작성하고 스크립트를 실행함
gedit install.sh

# install.sh gedit후 복사
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

sh install.sh

# 추가로 필요한 툴 
apt install vim net-tools -y
```

- Master Node 

```shell script
swapoff -a # 스왑기능 삭제 (현재 상태)
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab # 스왑기능삭제 (리부팅 후 부터 계속)


kubeadm init # 마스터노드로 사용

#마스터 노드에서만
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 끝에 조인할수 있도록 명령어를 만들어줌(워커 노드에서 실행, 관리자 권한으로 실행)
kubeadm join 10.211.55.5:6443 --token xc91kx.2xti8tceoeb1mito \
    --discovery-token-ca-cert-hash sha256:8ddc5d70f1fb6d98fca40e5234b7a8a10b9d1f2bc8dc4623e2954105a8d88158
```

- Slave Node

```shell script
swapoff -a # 스왑기능 삭제 (현재 상태)
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab # 스왑기능삭제 (리부팅 후 부터 계속)
kubeadm join 10.211.55.5:6443 --token xc91kx.2xti8tceoeb1mito \
    --discovery-token-ca-cert-hash sha256:8ddc5d70f1fb6d98fca40e5234b7a8a10b9d1f2bc8dc4623e2954105a8d88158
# weave 설치 및 등록
kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
```