FROM node:18-buster
WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y cmake && \
    rm -rf /var/lib/apt/lists/*

COPY package.json ./
    RUN npm config set registry https://registry.npmmirror.com
RUN npm install
COPY . .

EXPOSE 3000
CMD ["node" , "index.js"]
