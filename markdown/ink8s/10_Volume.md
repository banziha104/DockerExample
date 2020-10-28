# 볼륨 

- 컨테이너가 외부 스토리지에 엑세스하고 공유하는  방법
- 포드의 각 컨테이너에는 분리된 파일 시스템 존재
- 볼륨은 포드의 컴포넌트이며 포드의 스펙에 의해 정의
- 독립적인 쿠버네티스 오브젝트가 아니며, 스스로 생성, 삭제 불가
- 각 컨테이너의 파일 시스템의 볼륨을 마운트하여 생성 
- 볼륨 종류
    - 임시볼륨 
      - EmptyDir : 일시적인 데이터 저장
      ```yaml
      apiVersion: v1
      kind: Pod
      metadata:
      name: count
      spec:
      containers:
      - image: gasbugs/count
        name: html-generator
        volumeMounts:
        - name: html
          mountPath: /var/htdocs
      - image: httpd
        name: web-server
        volumeMounts:
        - name: html 
          mountPath: /usr/local/apache2/htdocs
          readOnly: true
          ports:
        - containerPort: 80
          protocol: TCP
          volumes:
      - name: html
        emptyDir: {}
      ```
      
    - 로컬볼륨 
      - hostpath : 포드에 호스트 노드의 파일 시스템에서 파일이나 디렉토리를 마운트
        ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
          name: hostpath-http
        spec:
          containers:
          - image: httpd
            name: web-server
            volumeMounts:
            - name: html
              mountPath: /usr/local/apache2/htdocs
              readOnly: true
            ports:
            - containerPort: 80
              protocol: TCP
          volumes:
          - name: html
            hostPath: # 호스트 패스
               path: /var/htdocs
               type: Directory
        ```
    - 네트워크 볼륨
      - nfs : 기존 NFS 공유가 포드에 장착
    - 클라우드 종속적 네트워크 볼륨 
      - gcePersistentDis : 구글 컴퓨트엔진 영구디스크 마운트
        ```shell 
        # gcloud 디스크 생성
        gcloud compute disks create --size=10GiB --zone=asia-northeast3-a mongodb
        ```
        ```yaml
        apiVersion: v1
        kind: Pod
        metadata:
          name: mongodb
        spec:
          containers:
            - image: mongo
              name: mongodb
              volumeMounts:
                - mountPath: /data/db
                  name: mongodb
          volumes:
            - name: mongodb
              gcePersistentDisk:
                pdName: mongodb
                fsType: ext4
        ```
      - awsEBS : 아마존 영구디스크 마운트 
      - azureFile : 애저 영구디스크 마운트
    

```shell
# NFS 서버 설치 
apt-get update
apt-get install nfs-common nfs-kernel-server portmap

# 디렉토리 생성
mkdir /home/nfs
chmod 777 /home/nfs

# 수정
vim /etc/exports  # /home/nfs 0.0.0.0을 적음

mount -t nfs <nfs서버 IP>:/home/nfs /mnt
```

<br>

# 볼륨 추상화


- PersistentVolume (PV)
- PersistentVolumeClaim (PVC)

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mongodb
spec:
  containers:
  - image: mongo
    name: mongodb
    volumeMounts:
    - mountPath: /data/db
      name: mongodb
  volumes:
  - name: mongodb
    persistentVolumeClaim:
      claimName: mongo-pvc
---
# PVC
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongo-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
  storageClassName: ""
---
# PV
apiVersion: v1
kind: PersistentVolume
metadata:
  name: mongo-pv
spec:
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  - ReadOnlyMany
  persistentVolumeReclaimPolicy: Retain # Retain(유지), Delete(외부 인프라와 연관된 스토리지 자산 모두 제거), Recycle(볼븀에 대한 스크럽을 수행하고 새 클라임에 대해 다시 사용할수 있음)
  gcePersistentDisk:
    pdName: mongodb
    fsType: ext4
```

- 동적 프로비저닝 : PV를 직접 만드는 대신 사용자가 원하는 PV 유형을 선택하도록 오브젝트 정의 가능.
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: storage
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mongo-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 10Gi
  storageClassName: storage
---
apiVersion: v1
kind: Pod
metadata:
  name: mongodb
spec:
  containers:
  - image: mongo
    name: mongodb
    volumeMounts:
    - mountPath: /data/db
      name: mongodb
  volumes:
  - name: mongodb
    persistentVolumeClaim:
      claimName: mongo-pvc
```

<br>

# 스테이트 풀셋 

> 어플리케이션의 상태를 저장하고 관리하는데 사용되는 쿠버네티스 객체.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
  - port: 80
    name: web
  clusterIP: None
  selector:
    app: nginx
```

```yaml

apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: web
spec:
  selector:
    matchLabels:
      app: nginx
  serviceName: "nginx" # 헤드레스 서비스를 지정한다.
  replicas: 3 # by default is 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      terminationGracePeriodSeconds: 10 # 강제 종료까지 대기하는 시간
      containers:
      - name: nginx
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates: # PVC 설정을 저장하는 부분
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      storageClassName: "standard"
      resources:
        requests:
          storage: 1Gi
```