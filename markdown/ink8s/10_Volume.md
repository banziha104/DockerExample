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
        ```yaml
        piVersion: v1
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
    