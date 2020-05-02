# 쿠버네티스 베이직 

<br>

### Kuberspray

> 클러스터 생성도구 

<br>

- 마스터노드 : 노드들의 상태를 관리하고 제어함. 
- 워커노드 : kubelet(프로세스,에이전트가) 동작하며, 마스터 노드의 명령을 받아 파드나 잡을 실행 
- 앤서블 기반 
- 앤서블을 이용한 원격 서버접근은 SSH로 이루어지기떄문에 모든 VM인스턴스에 SSH 키를 전송해야함 

### 마스터 노드 세팅

- SSH 키 생성과 배포
    - ssh키 만들기
    
    ```bash
    ssh-keygen -t rsa
    ```
    
    - ssh 키 확인 

    ```bash
    cat .ssh/id_rsa.pub
    ```
  
    - 구글 클라우드 플랫폼에 메타데이터 수정 
    - 다른 VM인스턴스들의 호스트네임과 내부 IP를 확인 
    ```bash
    ssh instance-2 hostname
    ```
<br>

- Kuberspary 설치
    - 파이썬 매니저인 pip 설치

    ```bash
    sudo apt update
    sudo apt -y install python-pip
    pip --version
    ```
  
    - Kuberspray 클론 및 설치 
    
    ```bash
    git clone https://github.com/kubernetes-sigs/kubespray.git
    cd kubespray
    git checkout -b v2.11.0
    ls -al
    cat requirements.txt 
    sudo pip install -r requirements.txt
    ```
    
    - 만약 로케일 오류가 날경우 
    
    ```bash
    export LC_ALL="en_US.UTF-8"
    export LC_CTYPE="en_US.UTF-8"
    sudo dpkg-reconfigure locales
    ```
    
    - kuberspray 설정 

    ```bash
    cp -rfp inventory/sample inventory/mycluster
    sudo apt install tree # 트리 추가 
    tree inventory/mycluster/group_vas # 디렉토리 구조확인
    vi inventory/mycluster/inventory.ini
    ```
  
    - inventory.ini를 다음과 같이 수정
    
    ```ini
    [all]
    # 클러스트로 구성될 서버들의 호스트네임과 IP를 설정함. 앤서블은 설정값으로 사용하는 IP가 같으면 호스트네임만 입력해 ssh로 통신가능
    instance-1 ansible_host=10.128.0.3 ip=10.128.0.3 etcd_member_name=etcd1
    instance-2 ansible_host=10.128.0.4 ip=10.128.0.4 etcd_member_name=etcd2
    instance-3 ansible_host=10.128.0.5 ip=10.128.0.5 etcd_member_name=etcd3
    instance-4 ansible_host=10.128.0.6 ip=10.128.0.6 
    instance-5 ansible_host=10.128.0.7 ip=10.128.0.7 
    
    # ## configure a bastion host if your nodes are not directly reachable
    # bastion ansible_host=x.x.x.x ansible_user=some_user
    # 마스터 노드로 사용할 서버의 호스트네임을 설정함. all에서 설정했기떄문에 이름만 제공
    [kube-master] 
    instance-1
    instance-2
    instance-3
  
    # 쿠버네티스 클러스터 데이터를 저장하는 etcd를 설치할 호스트네임을 지정. 마스터 노드를 별도로 구성할 수 있음
    [etcd] 
    instance-1
    instance-2
    instance-3
    
    # 워커노드로 사용할 서버의 호스트네임을 설정 
    [kube-node] 
    instance-4
    instance-5
    
    [calico-rr]
    
    # 쿠버네티스를 설치할 노드들을 설정, 보통 기본 설정 그대로 사용
    [k8s-cluster:children]
    kube-master
    kube-node
    calico-rr
    
    ``` 
  
    - 앤서블 가동 
    
    ```bash
    cd kuberspray
    ansible-playbook -i inventory/mycluster/inventory.ini -v --become --become-user=root cluster.yml
    ```
    
    - 설치 확인 
    ```bash
    sudo -i
    kubectl get node
    ```
            
