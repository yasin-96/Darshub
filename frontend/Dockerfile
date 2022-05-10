#==================== Building Stage ================================================
# base image
FROM node:12.2.0-alpine as node

# set working directory
#RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
#ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
ARG envname

COPY package*.json ./
RUN npm install
COPY . .

RUN echo "Will now run npm run build:env with environment: -$envname-"
RUN ENVNAME=${envname} npm run build:${envname}

#==================== Setting up stage ====================
FROM nginx
COPY --from=node /usr/src/app/dist /usr/share/nginx/html

# overriding default nginx config
#COPY --from=node /usr/src/app/conf/nginx.conf /etc/nginx/conf.d/default.conf
