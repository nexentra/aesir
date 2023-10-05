/** @type {import('next').NextConfig} */
const withNextra = require('nextra')({
    theme: 'nextra-theme-docs',
    themeConfig: './theme.config.tsx',
    experimental: {
      outputStandalone: true,
    },
  })
   
  module.exports = withNextra()