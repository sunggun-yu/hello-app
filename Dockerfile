FROM alpine

COPY hello-app /app/hello-app

EXPOSE 8080

CMD ["/app/hello-app"]
