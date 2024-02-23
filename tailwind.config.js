/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ}", "./**/*.templ"],
  safelist: [],
  plugins: [require("daisyui")],
  daisyui: {themes: ["light"]},
}