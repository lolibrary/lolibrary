FROM amelia/dhparam:latest as dhparam
FROM nginx:mainline-alpine as production

COPY --from=dhparam /etc/ssl/dhparam.pem /etc/nginx/dhparam.pem
COPY ./nginx.conf /etc/nginx/nginx.conf

HEALTHCHECK --interval=5s CMD "curl -f http://localhost:8080/status"
