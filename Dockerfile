FROM acoshift/go-scratch

COPY entrypoint /entrypoint

EXPOSE 8080
EXPOSE 9000

ENTRYPOINT ["/entrypoint"]
