module.exports = {
    transpileDependencies: [
        'vuetify'
    ],
    devServer: {
        compress: true,
        disableHostCheck: true, //webpack4.0 开启热更新
    }
}
