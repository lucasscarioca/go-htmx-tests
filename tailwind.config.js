/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/views/**/*.html',
    'assets/static/js/**/*.js'
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // Orange
        // 'accent-light': '#F9C8B5',
        // 'accent': '#EF8354',
        // 'accent-bold': '#E85000',
        // Green
        'accent-light': '#00dab4',
        'accent': '#00c193',
        'accent-bold': '#00a074',
        'cancel-light': '#F9A29F',
        'cancel': '#FE5F55',
        'cancel-bold': '#f64a3a',
        'secondary': '#F3F3F3', //#BFC0C0
        'secondary-dark': '#4F5D75',
        'primary': '#FAFAFA',
        'primary-dark': '#2D3142',
      },
      fontFamily: {
        primary: ['"Inter"'], //, 'ui-sans-serif', 'system-ui'
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  }
}
