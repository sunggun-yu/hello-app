FROM alpine

WORKDIR /app

COPY ./templates/ templates/
COPY hello-app hello-app

EXPOSE 8080

CMD ["/app/hello-app"]
