FROM ubuntu:focal

COPY simple /simple

ENTRYPOINT [ "/simple" ]
CMD ["--help"]
