# 使用官方Redis镜像作为基础镜像
FROM redis:latest

# 将Redis配置文件复制到容器中
COPY redis.conf /usr/local/etc/redis/redis.conf

# 将自定义的Redis持久化数据目录挂载到容器中
VOLUME /data

# 暴露Redis默认端口6379
EXPOSE 6379

# 启动Redis服务
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]

#docker run --name gin-project-gredis -d -p 6379:6379  -v D:/github_project/ginproject/project-common/docker/gredis/gredis-data:/var/lib/gredis gredis:latest