module.exports = {
  plugins: [
    'emotion'
  ],
  presets: [
    [    
      '@babel/preset-env',
      {
        targets: {
          node: 'current'
        }
      }
    ],
    '@babel/preset-typescript',
  ]
}
