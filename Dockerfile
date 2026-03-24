FROM alpine:3.23.3

WORKDIR /app

COPY ./templates/ templates/
COPY ./assets/ assets/
COPY hello-app hello-app

CMD ["/app/hello-app"]
