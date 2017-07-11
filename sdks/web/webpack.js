const path = require('path');

var config = {
    entry: {
        main: './src/ratings.js'
    },
    output: {
        filename: 'ratings.js',
        path: path.resolve('.')
    }
};

module.exports = config;