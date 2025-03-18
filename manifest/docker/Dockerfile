FROM loads/alpine:3.8 as backend-builder

ENV WORKDIR /project

ADD hack $WORKDIR/hack/
ADD manifest/config $WORKDIR/manifest/config/
ADD resource $WORKDIR/resource/
ADD ./temp/authz $WORKDIR/authz
ADD ./manifest/docker/cmd.sh $WORKDIR/cmd.sh
RUN chmod +x $WORKDIR/authz
RUN chmod +x $WORKDIR/cmd.sh

FROM node:20-slim as frontend-builder
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

ARG REACT_APP_API_BASE_URI=http://localhost:8080/v1
ARG REACT_APP_DEMO_ENABLED=false

ENV REACT_APP_API_BASE_URI=$REACT_APP_API_BASE_URI
ENV REACT_APP_DEMO_ENABLED=$REACT_APP_DEMO_ENABLED

WORKDIR /app

COPY ui/public ./public/
COPY ui/src ./src/
COPY ui/index.html .
COPY ui/package.json .
COPY ui/pnpm-lock.yaml .
COPY ui/tsconfig.json .
COPY ui/vite.config.js .

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build


FROM ghcr.io/nginxinc/nginx-unprivileged:1.27-alpine

ARG GRPC_ADDRESS=:8081
ARG SERVER_ADDRESS=:8080
ARG NGINX_VHOST=./manifest/docker/vhost.conf

ENV GRPC_ADDRESS=$GRPC_ADDRESS
ENV SERVER_ADDRESS=$SERVER_ADDRESS

COPY --from=backend-builder /project/authz /authz
COPY --from=backend-builder /project/cmd.sh /cmd.sh
COPY --from=frontend-builder /app/build /usr/share/nginx/html

COPY ${NGINX_VHOST} /etc/nginx/conf.d/default.conf

EXPOSE 3000 ${GRPC_ADDRESS} ${SERVER_ADDRESS}

ENTRYPOINT [ "/cmd.sh" ]
