'use strict';

const fs = require('fs');
const rollup = require('rollup');
const uglify = require('rollup-plugin-uglify');
const resolve = require('rollup-plugin-node-resolve');
const commonjs = require('rollup-plugin-commonjs');
const del = require('del');
const pkg = require('./package.json');

const resolvePlugin = resolve({
    // use "jsnext:main" if possible – see
    // https://github.com/rollup/rollup/wiki/jsnext:main
    jsnext: true, // Default: false

    // some package.json files have a `browser` field which specifies alternative
    // files to load for people bundling for the browser. If that's you, use this
    // option, otherwise pkg.browser will be ignored
    browser: true, // Default: false

    // not all files you want to resolve are .js files
    extensions: [
        '.js', '.json'
    ], // Default: ['.js']

    // whether to prefer built-in modules (e.g. `fs`, `path`) or local ones with the
    // same names
    preferBuiltins: false, // Default: true
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

promise = promise.then(() => del(['dist/*']));

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

promise = promise.then(() => {
    delete pkg.devDependencies;
    delete pkg.scripts;
    delete pkg.eslintConfig;

    fs.writeFileSync('dist/package.json', JSON.stringify(pkg, null, '  '), 'utf-8');
    fs.writeFileSync('dist/LICENSE.md', fs.readFileSync('LICENSE.md', 'utf-8'), 'utf-8');
});

promise.catch(err => console.error(err.stack)); // eslint-disable-line no-console