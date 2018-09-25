FROM acoshift/go-scratch

COPY entrypoint /entrypoint

COPY config /config

EXPOSE 8080

ENTRYPOINT ["/entrypoint"]
