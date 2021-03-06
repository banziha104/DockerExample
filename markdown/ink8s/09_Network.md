# 네트워크 

- 한 포드에 있는 다수의 컨테이너끼리 통신
    - 인터페이스를 공유하는 아무 동작을 하지 않는 pause 컨테이너를 생성해 통신
    - 포트를 겹치게 구성하지 못하는 것이 특징
    - 각 포드마다 하나의 pause 이미지 실행 
    - ![network1](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/image/network1.png)
- 포드끼리 통신
    - 포드끼리의 통신을 위해서는 CNI 플러그인이 필요
    - LAN을 활용해 통신
    - 다른 노드에 있는 팟과 통신해야할땐 Weave DaemonSet을 이용(Weave의 경우)
    ```shell script
    sudo netstat -antp | grep weave
    tcp        0      0 127.0.0.1:6784          0.0.0.0:*               LISTEN      5744/weaver         
    tcp        0      0 10.211.55.5:57534       10.96.0.1:443           ESTABLISHED 4586/weave-npc      
    tcp6       0      0 :::6781                 :::*                    LISTEN      4586/weave-npc      
    tcp6       0      0 :::6782                 :::*                    LISTEN      5744/weaver         
    tcp6       0      0 :::6783                 :::*                    LISTEN      5744/weaver         
    tcp6       0      0 10.211.55.5:6783        10.211.55.7:42777       ESTABLISHED 5744/weaver        # 다른 노드와 연결  
    tcp6       0      0 10.211.55.5:6783        10.211.55.6:42547       ESTABLISHED 5744/weaver        # 다른 노드와 연결
    ``` 
    - ![network2](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/image/network2.png)
    - ![network3](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/image/network3.png)
- 포드와 서비스 사이의 통신
    - ClusterIP를 생성하면 iptables를 생성
    - kube-proxy라는 컴포넌트로 서비스 트래픽을 제어
    - iptables는 리눅스 커널 기능인 netfilter를 사용하여 트래픽을 제어
- 외부 클라이언트와 서비스 사이의 통신
    - netfilter와 kube-proxy 기능을 사용해 원하는 서비스 및 포드로 연결 
    - ![network5](https://github.com/banziha104/DockerExample/blob/master/markdown/ink8s/image/network4.png)
    
<br>

# CoreDNS

> 서비스를 생성하면, 대응되는 DNS 엔트리가 생성

- <서비스 이름>.<네임스페이스-이름>.svc.cluster.local의 형식을 가짐
- 내부에서 DNS 서버 역할을 하는 POD이 존재
- 각 미들웨어를 통해 로깅,캐시,쿠버네티스를 질의하는 등의 기능을 가짐.
- 해당 DNS에는 configmap 저장소를 사용해 설정 파일을 컨트롤 
- POD에도 서브도메인을 활용해 도메인 사용가능.
