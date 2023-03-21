FROM node:lts-alpine

# устанавливаем простой HTTP-сервер для статики
RUN npm install -g http-server

WORKDIR /app

# копируем оба 'package.json' и 'package-lock.json'
COPY package*.json ./

RUN npm install

COPY . .

# собираем приложение
RUN npm run build

EXPOSE 8080
CMD [ "http-server", "dist" ]