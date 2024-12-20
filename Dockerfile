FROM alpine

WORKDIR /app

COPY ./templates/ templates/
COPY ./assets/ assets/
COPY hello-app hello-app

CMD ["/app/hello-app"]
