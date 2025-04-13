FROM node:22.14-bookworm

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 3000

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]