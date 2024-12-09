FROM cgr.dev/chainguard/static:latest
USER 65532:65532
WORKDIR /app
COPY --chown=65532:65532 ./bin/api /app/
EXPOSE 8080
CMD [ "/app/api" ]
