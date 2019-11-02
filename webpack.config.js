const path = require('path');

module.exports = {
    entry: './websrc/js/index.js',
    output: {
        filename: 'main.js',
        path: path.resolve(__dirname, 'webpublic/js'),
    },
    mode: process.env.ENV === 'production' ? 'production' : 'development',
    module: {
        rules: [
            {
                test: /\.js$/,
                // include: [
                //     path.resolve(__dirname, "app")
                // ],
                enforce: "pre",
                enforce: "post",
                loader: "babel-loader",
                options: {
                    presets: ["@babel/env"]
                },
              },
        ]
    },
};