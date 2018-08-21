FROM scratch

ADD ./web/static /web/static/
COPY main.out /

EXPOSE 8220
ENTRYPOINT ["/main.out"]