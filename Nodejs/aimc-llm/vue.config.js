const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  publicPath: './',
  transpileDependencies: true,
  lintOnSave: false,
  devServer: {
    proxy: process.env.VUE_APP_BASE_API
  }
})
