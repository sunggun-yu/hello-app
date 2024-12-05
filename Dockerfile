FROM alpine

WORKDIR /app

COPY ./templates/ templates/
COPY hello-app hello-app

CMD ["/app/hello-app"]
