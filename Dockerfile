FROM cgr.dev/chainguard/static:latest

WORKDIR /app

COPY ./bin/api /app/

EXPOSE 8080

CMD [ "/app/api" ]
