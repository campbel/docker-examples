FROM nginx:alpine
RUN apk update && apk add gettext
ADD nginx.conf.template /etc/nginx/nginx.conf.template
ENV message "hello, environment substitute!"
CMD envsubst < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf && cat /etc/nginx/nginx.conf && nginx -g 'daemon off;'
