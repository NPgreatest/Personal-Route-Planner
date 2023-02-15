# Personal-Route-Planner
A route planner system supported site review, comment sites and personal route planning 
developed by Golang and Vue
# 1. The composition of this program
* back-end: Golang
* front-end:Vue.js
* deploy:Docker-compose

### The framework of back-end:
> sqlx, gin, json-web-token, viper, snowflake, base64-encoding-decoding

* The database of back-end:Mysql
* All pictures are stored in /images in the static way, and access in website 
in https://127.0.0.1/images/example.jpg


### The framework of front-end:
> element-ui, e-chats, animation


# 2.Get started
Before starting, you should already install golang, docker and nodejs in your develop env.

you can run 
```docker-compose up``` . or run back-end and front-end separately.

Run back-end separately in GoLand

Run front-end separately by
```
cd vue
npm i
npm run serve
```
Then you can access website in http://localhost:8080/

# 3. TO-DO LIST
- [x] recommend algorithm
- [ ] connect with 高德API
- [ ] connect with chatgpt-ai
- [ ] redis server
- [ ] change to GORM framework



