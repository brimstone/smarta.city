FROM scratch

EXPOSE 8080

ENV PORT=8080

COPY app /

ENTRYPOINT ["/app"]

CMD ""
