FROM scratch

ENV PORT 8000
ENV MYSQL_DATABASE root:root@tcp(docker.for.mac.localhost:3306)/sme?charset=utf8mb4&parseTime=true
EXPOSE $PORT

COPY cms /
CMD [ "/cms" ]