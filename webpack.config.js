const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const TerserPlugin = require('terser-webpack-plugin');
const CssMinimizerPlugin = require('css-minimizer-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

module.exports = {
    entry: './resource/js/app.js',
    // output setting
    output: {
        filename: 'js/app.js', // js output
        path: path.join(__dirname, 'static') // output path
    },
    module: {
        rules: [{
                test: /\.html$/,
                loader: "html-loader"
            },
            {
                // extension .js case
                test: /\.js$/,
                use: [{
                    // use babel
                    loader: 'babel-loader',
                    // babel options
                    options: {
                        presets: [
                            // es2020 to es5
                            '@babel/preset-env',
                        ]
                    }
                }]
            },
            {
                test: /\.svg$/,
                use: [{
                    loader: 'html-loader',
                    options: {
                        minimize: true
                    }
                }]
            },
            {
                test: /\.(sa|sc|c)ss$/,
                exclude: /node_modules/,
                use: [
                    MiniCssExtractPlugin.loader,
                    {
                        loader: 'css-loader',
                        options: { url: false }
                    },
                    'sass-loader'
                ]
            },
        ]
    },
    plugins: [
        new CleanWebpackPlugin(),
        new MiniCssExtractPlugin({
            filename: 'css/app.css', // css output
        }),
    ],
    optimization: {
        minimize: true,
        minimizer: [
            new TerserPlugin(), // js minify
            new CssMinimizerPlugin() // css minify
        ]
    }
};