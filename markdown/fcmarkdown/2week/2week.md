# Docker 

- docker run -d -p 8080:8080 -p 5000:5000 --name myjenkins jenkins : 두 개의 포트를씀
- docker run -p 80:80 -v /home/docker/nginx.conf:/etc/nginx/nginx.conf --link myjenkins:jenkins -d nginx