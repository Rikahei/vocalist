module.exports = {
  devServer: {
    proxy: 'http://backend-go:1323'
  },
  "transpileDependencies": [
    "vuetify"
  ]
}