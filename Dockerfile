FROM scratch

COPY simple /simple

ENTRYPOINT [ "/simple" ]
CMD ["--help"]
