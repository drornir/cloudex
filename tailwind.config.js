/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pkg/component/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    require("@tailwindcss/aspect-ratio"),
  ],
};
