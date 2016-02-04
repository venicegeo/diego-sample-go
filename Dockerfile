FROM klokantech/gdal

ENTRYPOINT ["/diego-sample-go"]

COPY diego-sample-go /diego-sample-go
RUN chmod a+x /diego-sample-go
