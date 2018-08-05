# 멀티 스테이지 빌드

> 도커 이미지 빌드전에 어플리케이션을 빌드 해야할 때

```dockerfile
# 시작 스테이지
# build-env로 변수명을만듬
FROM maven:3.5.0-jdk-8-apline AS build-env
ADD ./pom.xml pom.xml
ADD ./src src/
RUN mvn clean package


# 끝 스테이지

FROM FROM openjdk:9-jre
# 위 스테이지를 가져옮
COPY --from=build-env target/app.jar app.jar
RUN java -jar app.jar

```


```dockerfile
FROM golang:1.7.3
WORKDIR /go/src/github.comn/alexellis/href-counter/
RUN go get -d -v golang.org/x/net/html  
COPY app.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/alexellis/href-counter/app .
CMD [“./app”]
```