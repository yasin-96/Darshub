#==================== Building Stage ================================================
# base image
FROM node:16.18.0-alpine as node

# clean working directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

ARG envname

RUN apk add --no-cache git

COPY package*.json ./
RUN npm i
RUN npm i --save-dev

COPY . .

RUN echo "Will now run npm run build:env with environment: -$envname-"
RUN ENVNAME=${envname} npm run build:${envname}

#==================== Setting up stage ====================
FROM nginx
COPY --from=node /usr/src/app/dist /usr/share/nginx/html

# overriding default nginx config
#COPY --from=node /usr/src/app/conf/nginx.conf /etc/nginx/conf.d/default.conf
