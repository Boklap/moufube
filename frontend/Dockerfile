FROM node:24.9.0-alpine

RUN apk add --no-cache python3 make g++ bash

ENV PATH="/usr/local/bin:$PATH"

WORKDIR /app
COPY package*.json ./

RUN npm ci

ENV PATH="/app/node_modules/.bin:$PATH"

COPY . .

CMD npm run dev