const webpack = require('webpack');
const path = require('path');
module.exports = {
    mode:"development",
  entry: './src/index.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, 'dist'),
  },
  target: 'node',
  module:{
    rules:[
        {
            test:/\.ts/,
            exclude: /node_modules/,
            use: {
                loader: 'ts-loader'
           }
        }
    ]
  },
  resolve:{
    extensions:['.ts', '.js']
  }
};