FROM scratch

EXPOSE 8080

ENV PORT=8080

COPY app /

COPY static /static

ENTRYPOINT ["/app"]

CMD ""
