FROM quay.io/PolyglotSystems/golang-ubi:latest AS build

WORKDIR /opt/app-root/src
COPY . .
RUN make build

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest AS bin

COPY --from=build /opt/app-root/src/dist/go-cart /opt/app-root/bin/
COPY container_root/ /

RUN mkdir -p /opt/app-root/go-cart/ \
 && chmod 777 /opt/app-root/go-cart/

EXPOSE 8080

CMD [ "/opt/app-root/bin/container_start.sh" ]