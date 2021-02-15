import commonjs from '@rollup/plugin-commonjs'
import babel from 'rollup-plugin-babel'
import replace from 'rollup-plugin-replace'
import html from 'rollup-plugin-html2'
import livereload from 'rollup-plugin-livereload'
import resolve from '@rollup/plugin-node-resolve'
import serve from 'rollup-plugin-serve'
import svelte from 'rollup-plugin-svelte'
import { terser } from 'rollup-plugin-terser'
import typescript from '@rollup/plugin-typescript'
import preprocess from 'svelte-preprocess'
import postcss from 'rollup-plugin-postcss'
import fs from 'fs'

const isDev = process.env.NODE_ENV === 'development'
const host = 'budget-holder.ru'
const port = 3000

// define all our plugins
const plugins = [
  replace({
    'process.env.NODE_ENV': JSON.stringify('development')
  }),
  svelte({
    dev: isDev,
    extensions: ['.svelte'],
    preprocess: preprocess()
  }),
  resolve({
    browser: true,
    dedupe: ['svelte']
  }),
  typescript(),
  babel({
    exclude: 'node_modules/**'
  }),
  postcss({
    plugins: []
  }),
  commonjs(),
  html({
    template: 'src/index.html',
    fileName: 'index.html'
  })
]

if (isDev) {
  plugins.push(
    // like a webpack-dev-server
    serve({
      open: true,
      contentBase: './dist',
      historyApiFallback: true, // for SPAs
      port,
      // host,
      https: {
        key: fs.readFileSync('./https/privateKey.key'),
        cert: fs.readFileSync('./https/certificate.crt')
      },
      onListening: function (server) {
        const address = server.getAddress()
        const host = address.host === '::' ? 'localhost' : address.host
        // by using a bound function, we can access options as `this`
        const protocol = this.https ? 'https' : 'http'
        console.log(`Server listening at ${protocol}://${host}:${address.port}/`)
      }
    }),
    livereload({ watch: './dist' })
  )
} else {
  plugins.push(terser({ sourcemap: isDev }))
}

module.exports = {
  input: 'src/index.js',
  output: {
    name: 'bundle',
    file: 'dist/bundle.js',
    sourcemap: isDev,
    format: 'iife'
  },
  plugins
}
