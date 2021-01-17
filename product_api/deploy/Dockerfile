FROM node:lts-alpine
WORKDIR /usr/src/app
COPY . .
COPY config/default.json.sample ./config/default.json
RUN npm install --ignore-engines
EXPOSE 3030
CMD ["npm", "start"]