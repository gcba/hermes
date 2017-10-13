'use strict';

const fs = require('fs');
const rollup = require('rollup');
const uglify = require('rollup-plugin-uglify');
const uglifyES = require('uglify-es');
const babel = require('rollup-plugin-babel');
const nodeResolve = require('rollup-plugin-node-resolve');
const commonjs = require('rollup-plugin-commonjs');
const del = require('del');

const resolvePlugin = nodeResolve({
    jsnext: true, // Default: false
    browser: true, // Default: false
    extensions: [
        '.js', '.json'
    ], // Default: ['.js']
    preferBuiltins: false, // Default: true
});

const babelPlugin = babel({
    include: 'src/**',
    exclude: 'node_modules/**',
    presets: [['es2015', {'modules': false}], 'stage-3'],
    runtimeHelpers: true,
    babelrc: false
});

const uglifyPlugin = uglify({}, uglifyES.minify);

const bundles = [
    {
        format: 'umd',
        ext: '.js',
        plugins: [babelPlugin],
        moduleName: 'Ratings'
    }, {
        format: 'umd',
        ext: '.min.js',
        plugins: [babelPlugin, uglifyPlugin],
        moduleName: 'Ratings',
        minify: true
    }
];

let promise = Promise.resolve();

promise = promise.then(() => del(['dist/*']));

for (const config of bundles) {
    let rollupConfig = {
        entry: 'src/main.js',
        plugins: [resolvePlugin, commonjs()].concat(config.plugins),
        onwarn: (warning) => {
            if (warning.code === 'THIS_IS_UNDEFINED') return;
            if (warning.message) console.warn(warning.message);
        }
    };

    let bundleConfig = {
        dest: `dist/ratings${config.ext}`,
        format: config.format,
        sourceMap: !config.minify,
        moduleName: config.moduleName
    };

    promise = promise.then(() => rollup.rollup(rollupConfig).then(bundle => bundle.write(bundleConfig)));

    if (config.ext === '.min.js') {
        let exampleConfig = {
            dest: `example/js/ratings${config.ext}`,
            format: bundleConfig.format,
            sourceMap: false,
            moduleName: bundleConfig.moduleName
        };

        promise = promise.then(() => rollup.rollup(rollupConfig).then(bundle => bundle.write(exampleConfig)));
    }
}

promise.catch(err => console.error(err.stack)); // eslint-disable-line no-console