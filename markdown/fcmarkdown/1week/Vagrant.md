# Vagrant

> 다양한 가상머신을 동일한 방식으로 만들고 관리할 수 있게 해주는 도구
> 루비 언어를 따르는 Vagrantfile 에 필요한 설정을 정의

- Vagrant Box : Vagrant 로 만들 수 있는 가상서버 유형
- Vagrantfile에 포함되는 내용
    - Box : 패키징할 서버 유형
    - 포트 바인딩 
    - 네트워크 주소
    - 이름,CPU,Memory
    - 부트 스트래핑 : 어플리케이션 설치
    - 프로비저닝 : 어플리케이션 설정
    - 폴더 마운트 
    
    
### CoreOS

> Docker 호스트에 특화된 리눅스
> 다만 패키지매니저는 없고, Docker로 대신함

- 최소화된 메모리 사용
- systemd(리눅스 구동시 서비스를 자동으로 뜨게 해), etcd(Key-value 스토어), fleet
- CoreUpdate(상용)
- https://coreos.com/


- fleet

[fleet](https://github.com/banziha104/DockerExample/img/2week/fleet.png)

### 실습

1. git clone https://github.com/coreos/coreos-vagrant.git : CoreOS 설치를 위한 Vagrantfile 을 받음
2. vagrant : vagrant 실행
3. user-data 파일과 config.rb 두 가지만 고침
4. cp config.rb.sample config.rb : 복사
5. 수정

```ruby
# Size of the CoreOS cluster created by Vagrant
$num_instances=1

# Used to fetch a new discovery token for a cluster of size $num_instances
$new_discovery_url="https://discovery.etcd.io/new?size=#{$num_instances}"

# Automatically replace the discovery token on 'vagrant up'

if File.exists?('user-data') && ARGV[0].eql?('up')
  require 'open-uri'
  require 'yaml'

  token = open($new_discovery_url).read

"config.rb" 80L, 2971C

```

```ruby

$num_instacnes = 1

$update_channel : ' stable '

$forwarded_ports = { 49153 => 49153 }

$shared_folders = { '/apps' => '/apps' }

$new_discovery_url = 'https://discovery.etcd.io/new' 

```

6. user-data 수정
7. vagrant up : 머신 실행
8. vagrant status : 머신 확인
9. vagrant ssh core-01 : 머신 접속


