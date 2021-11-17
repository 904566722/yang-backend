FROM centos-make-go:8.2.2004
RUN mkdir yang/
COPY . yang/
RUN cd yang/ && make compile-yang

FROM centos:8.2.2004
COPY --from=0 /yang/_out/bin/yang yang
RUN chmod a+x ./yang
CMD ["./yang"]