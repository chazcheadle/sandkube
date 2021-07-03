FROM node:13-alpine3.11
WORKDIR /app
COPY package*.json ./
RUN npm install
ADD . .
CMD node index.js
EXPOSE 3000