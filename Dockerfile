FROM scratch

ENV SL_MODE=prod
ENV SL_BASE_PATH=/app
ENV SL_WEB_SERVER_ADDRESS=":9000"

ADD ./web/static /app/web/static/
COPY main.out /app
COPY ladrawex.yml /app
COPY ladrawex-prod.yml /app

EXPOSE 9000
ENTRYPOINT ["/app/main.out"]