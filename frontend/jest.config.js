const sveltePreprocess = require('svelte-preprocess')

module.exports = {
  transform: {
    '^.+\\.svelte$': ['jest-transform-svelte', { preprocess: sveltePreprocess() }],
    '^.+\\.(js?|ts?)$': 'babel-jest'
  },
  testRegex: '(.*)_test.(js?|ts?|svelte?)$',
  moduleFileExtensions: ['js', 'ts', 'svelte'],
  testPathIgnorePatterns: ['node_modules'],
  collectCoverage: false,
  coverageDirectory: '<rootDir>/test_coverage',
  collectCoverageFrom: [
    '**/*.{ts,svelte}',
    '!**/node_modules/**'
  ],
  coverageReporters: [
    'json',
    'html',
    'text-summary'
  ],
  bail: false,
  verbose: true,
  transformIgnorePatterns: ['node_modules'],
  setupFilesAfterEnv: ['@testing-library/jest-dom/extend-expect', 'jest-extended']
}
