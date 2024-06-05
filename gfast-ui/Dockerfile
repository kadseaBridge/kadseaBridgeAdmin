FROM node:10.24.1
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 9200
CMD ["npm", "run", "dev"]
