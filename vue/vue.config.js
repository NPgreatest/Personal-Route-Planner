
module.exports = {
  devServer: {              //设置本地域名
    host: "localhost",
    port: 8082,
    proxy: {          //设置代理解决跨域问题
      "/api": {
        target: "http://localhost:8080",    //要跨域的域名
        changeOrigin: true,                  //是否开启跨域
        ws:false
      },
    },
  },

   // publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
  publicPath : './' ,
  transpileDependencies: true,
  lintOnSave: false,
  assetsDir:'static',
  parallel:false,
}


