FROM node:lts-alpine

COPY package.json ./
COPY yarn.lock ./
RUN yarn install
RUN yarn global add @vue/cli

EXPOSE 8080
