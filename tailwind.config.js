/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./cmd/web/**/*.{templ, templ.go}'],
  safelist: ['animate-pulse'],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ['light', 'dark'],
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
};
