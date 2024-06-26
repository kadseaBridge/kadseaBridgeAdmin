FROM node:10.24.1
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 9200
CMD ["npm", "run", "dev"]

## 使用多阶段构建，首先从 Node.js 镜像开始构建阶段
#FROM node:10.24.1 as build-stage
#
## 设置工作目录
#WORKDIR /app
#
## 复制 package.json 和 package-lock.json
#COPY package*.json ./
#
## 安装依赖
#RUN npm install
#
## 复制所有文件到工作目录
#COPY . .
#
## 构建项目
#RUN npm run build:prod
#
## 使用轻量级 Nginx 镜像进行部署
#FROM nginx:alpine
#
## 复制构建后的文件到 Nginx 的静态文件目录
#COPY --from=build-stage /app/dist /usr/share/nginx/html
#
## 复制自定义的 Nginx 配置文件（如果有的话）
#COPY nginx.conf /etc/nginx/nginx.conf
#
## 暴露端口
#EXPOSE 9200
#
## 启动 Nginx
#CMD ["nginx", "-g", "daemon off;"]
#
