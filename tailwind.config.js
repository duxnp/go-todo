/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./cmd/web/**/*.{templ, templ.go}'],
  theme: {
    extend: {},
  },
  plugins: [require('@tailwindcss/typography'), require('daisyui')],
};
