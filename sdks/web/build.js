'use strict';

const fs = require('fs');
const rollup = require('rollup');
const uglify = require('rollup-plugin-uglify');
const resolve = require('rollup-plugin-node-resolve');
const commonjs = require('rollup-plugin-commonjs');
const pkg = require('./package.json');

let resolvePlugin = resolve({
    // use "module" field for ES6 module if possible
    module: true, // Default: true

    // use "jsnext:main" if possible – see
    // https://github.com/rollup/rollup/wiki/jsnext:main
    jsnext: true, // Default: false

    // use "main" field or index.js, even if it's not an ES6 module (needs to be
    // converted from CommonJS to ES6 – see
    // https://github.com/rollup/rollup-plugin-commonjs
    main: false, // Default: true

    // some package.json files have a `browser` field which specifies alternative
    // files to load for people bundling for the browser. If that's you, use this
    // option, otherwise pkg.browser will be ignored
    browser: false, // Default: false

    // not all files you want to resolve are .js files
    extensions: [
        '.js', '.json'
    ], // Default: ['.js']

    // whether to prefer built-in modules (e.g. `fs`, `path`) or local ones with the
    // same names
    preferBuiltins: false, // Default: true

    // Lock the module search in this path (like a chroot). Module defined outside
    // this path will be mark has external
    jail: '/', // Default: '/'

    // If true, inspect resolved files to check that they are ES2015 modules
    modulesOnly: false, // Default: false
});

const bundles = [
    {
        format: 'es',
        ext: '.mjs',
        plugins: []
    }, {
        format: 'cjs',
        ext: '.browser.js',
        plugins: []
    }, {
        format: 'umd',
        ext: '.js',
        plugins: [],
        moduleName: 'ratings'
    }, {
        format: 'umd',
        ext: '.min.js',
        plugins: [uglify()],
        moduleName: 'ratings',
        minify: true
    }
];

let promise = Promise.resolve();

for (const config of bundles) {
    promise = promise.then(() => rollup.rollup({
        entry: 'src/main.js',
        plugins: [resolvePlugin, commonjs()].concat(config.plugins)
    }).then(bundle => bundle.write({
        dest: `dist/${config.moduleName || 'ratings'}${config.ext}`,
        format: config.format,
        sourceMap: !config.minify,
        moduleName: config.moduleName
    })));
}

promise.catch(err => console.error(err.stack)); // eslint-disable-line no-console