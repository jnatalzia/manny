const CopyPlugin = require('copy-webpack-plugin');
const path = require('path');

module.exports = {
    entry: './websrc/js/index.js',
    output: {
        filename: 'main.js',
        path: path.resolve(__dirname, 'webpublic/js')
    },
    mode: process.env.ENV === 'production' ? 'production' : 'development',
    module: {
        rules: [
            {
                test: /\.js$/,
                enforce: 'pre',
                enforce: 'post',
                loader: 'babel-loader',
                options: {
                    presets: ['@babel/env']
                }
            },
            {
                test: /\.scss$/,
                use: ['style-loader', 'css-loader', 'sass-loader']
            }
        ]
    },
    plugins: [new CopyPlugin([{ from: 'websrc/views', to: '../' }])]
};
