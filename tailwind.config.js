/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./views/**/*.{templ,go}"],
  darkMode: "media",
  theme: {
    extend: {},
  },
  plugins: [
    require("tailwindcss-animated"),
    require("@tailwindcss/forms"),
  ],
}

