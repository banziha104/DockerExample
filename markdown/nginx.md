# nginx 

```shell
docker run --name nginx \
-d -p 8080:80 \
-v $(PWD)/sites-enabled:/etc/nginx/sites-enabled \
-v $(PWD)/log/nginx:/var/log/nginx \
saltfactory/nginx
```