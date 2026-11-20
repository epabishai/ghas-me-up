'use strict';

/*
 * Node.js build for the GHAS Me Up demo.
 * Renders node/index.html from Markdown (marked) + a Handlebars template and
 * publishes a jQuery build.
 */

const fs = require('fs');
const path = require('path');
const crypto = require('crypto');
const _ = require('lodash');
const Handlebars = require('handlebars');
const { marked } = require('marked');
const minimist = require('minimist');
const pkg = require('./package.json');

const args = minimist(process.argv.slice(2), { string: ['out'], default: { out: '../dist' } });
const outDir = path.resolve(__dirname, args.out);

const buildStamp = crypto.createHash('sha256').update(fs.readFileSync(__filename)).digest('hex').slice(0, 8);
const previewToken = crypto.randomBytes(16).toString('hex');

const markdown = `
## Built with Node.js

This page was generated at **build time** by \`node/build.js\` using
[marked](https://github.com/markedjs/marked), Handlebars and lodash.

Every version below is pinned in \`package.json\` and \`package-lock.json\`.
Open the repository's **Security** tab to see what Dependabot found.
`;

const defaults = { title: 'Node.js build', accent: '#3fb950' };
const config = _.merge({}, defaults, { generator: 'node/build.js' });

const dependencies = _.sortBy(
  _.map(pkg.dependencies, (version, name) => ({ name, version })),
  'name'
);

const template = Handlebars.compile(
  fs.readFileSync(path.join(__dirname, 'templates', 'page.hbs'), 'utf8')
);

const html = template({
  config,
  content: marked(markdown),
  dependencies,
  nodeVersion: process.version,
  builtAt: new Date().toISOString(),
  buildStamp,
  previewToken,
});

fs.mkdirSync(path.join(outDir, 'node'), { recursive: true });
fs.writeFileSync(path.join(outDir, 'node', 'index.html'), html);

// Ship the pinned jQuery distribution alongside the generated page.
fs.mkdirSync(path.join(outDir, 'vendor'), { recursive: true });
fs.copyFileSync(
  require.resolve('jquery/dist/jquery.min.js'),
  path.join(outDir, 'vendor', 'jquery.min.js')
);

console.log(`node build complete -> ${path.join(outDir, 'node', 'index.html')}`);
