FROM centos:8.2.2004
RUN mkdir yang/
COPY . yang/
RUN cd yang/ && make compile-controller
COPY /yang/_out/bin/yang yang
CMD ["./yang"]