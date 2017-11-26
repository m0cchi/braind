const ExtractTextPlugin = require('extract-text-webpack-plugin');
const webpack = require('webpack');
const path = require('path');

const index = {
  entry: __dirname + '/js/index.js',
  // 出力の設定
  output: {
    // 出力するファイル名
    filename: 'index.js',
    // 出力先のパス（v2系以降は絶対パスを指定する必要がある）
    path: path.join(__dirname, 'resources/public/js/')
  },
  plugins: [
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      'window.jQuery': 'jquery'
    })
  ]
};

const bootstrap = {
  entry:
    __dirname + '/node_modules/bootstrap/dist/css/bootstrap.min.css'
  ,
  output: {
      path: __dirname + '/resources/public/css/',
      filename: 'bootstrap.css'
  },
  module: {
    loaders: [
      {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract("css-loader")
      },
      {
        test: /\.(woff|woff2|eot|ttf|svg)$/,
        loader: 'file-loader?name=../font/[name].[ext]'
      }
    ]
  },
  plugins: [
    new ExtractTextPlugin('../css/bootstrap.css')
  ]
};


const header = {
  entry: __dirname + '/js/header.js',
  output: {
    filename: 'header.js',
    path: path.join(__dirname, 'resources/public/js/')
  },
  plugins: [
    new webpack.ProvidePlugin({
      $: 'jquery',
      jQuery: 'jquery',
      'window.jQuery': 'jquery',
      'Tether': 'tether'
    })
  ]
};


module.exports = [header, index, bootstrap]
