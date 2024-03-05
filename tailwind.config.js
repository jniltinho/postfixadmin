/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ}", "./**/*.templ"],
  safelist: [],
  plugins: [require("daisyui")],
  //daisyui: {themes: ["light"]},
  daisyui: {
    themes: [
      {
        light: {
          "primary": "#a991f7",
          "secondary": "#f6d860",
          "accent": "#37cdbe",
          "neutral": "#3d4451",
          "base-100": "#ffffff",

          "--rounded-box": "0", // border radius rounded-box utility class, used in card and other large boxes
          "--rounded-btn": "0", // border radius rounded-btn utility class, used in buttons and similar element
          "--rounded-badge": "0", // border radius rounded-badge utility class, used in badges and similar
          "--tab-border": "0", // border width of tabs
          "--tab-radius": "0", // border radius of tabs
        },
      },
    ],
  },
}