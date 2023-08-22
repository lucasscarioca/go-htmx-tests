/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/views/**/*.html',
  ],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily: {
        mono: ['Poppins', 'monospace'],
      }
    },
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  }
}
