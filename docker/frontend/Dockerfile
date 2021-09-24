ARG NODE_VERSION=alpine

FROM node:${NODE_VERSION}

LABEL maintainer="Neekity"

VOLUME /data

WORKDIR /data

ARG NPM_REGISTRY
ENV NPM_REGISTRY ${NPM_REGISTRY}

ARG INSTALL_NPM_GULP=false
ARG INSTALL_NPM_BOWER=false
ARG INSTALL_NPM_VUE_CLI=false
ARG INSTALL_NPM_ANGULAR_CLI=false

RUN if [ ${NPM_REGISTRY} ]; then \
        npm config set registry ${NPM_REGISTRY} \
    ;fi \
    && if [ ${INSTALL_NPM_GULP} = true ]; then \
        npm install -g gulp \
    ;fi \
    && if [ ${INSTALL_NPM_BOWER} = true ]; then \
        npm install -g bower \
    ;fi \
    && if [ ${INSTALL_NPM_VUE_CLI} = true ]; then \
     npm install -g @vue/cli \
    ;fi \
    && if [ ${INSTALL_NPM_ANGULAR_CLI} = true ]; then \
        npm install -g @angular/cli \
    ;fi

EXPOSE 8881

CMD ["/bin/sh", "-c", "yarn install && yarn run serve"]