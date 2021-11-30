module.exports = {
    outputDir: 'dist',   //build输出目录
    assetsDir: 'assets', //静态资源目录（js, css, img）
    lintOnSave: false, //是否开启eslint
    devServer: {
        // host: "10.25.16.212",
        host: "10.25.17.18",
        port: '8080',
        https: false,
        hotOnly: false,
        proxy: {
            '/week': {
                target: 'http://10.25.16.212:8981/',
                ws: true,
                changeOrigin: true,
                pathRewrite: {
                    '^/week': ''
                }
            }
        },
    },
    runtimeCompiler: true,
};

