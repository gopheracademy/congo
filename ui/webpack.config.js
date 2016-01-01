var path = require('path');

module.exports = {
  entry: path.resolve(__dirname, "app/Main.js"),
  devtool: "source-map",
  output: {
    path: path.resolve(__dirname, '..', "public"),
    filename: "app.js",
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel",
        query:
          {
            presets:["es2015","react"]
          }
      }
    ]
  },
};
