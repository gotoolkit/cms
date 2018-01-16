FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENV PORT 8000
ENV MYSQL_DATABASE root:root@tcp(docker.for.mac.localhost:3306)/sme?charset=utf8mb4&parseTime=true
ENV TELEGRAM_HORN_URL https://bot.wilead.com/api/tom
EXPOSE $PORT

COPY cms /
CMD [ "/cms" ]