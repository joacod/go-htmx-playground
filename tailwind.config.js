/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.{html,templ}", "./api/**/*.{html,templ}"],
  theme: {
    extend: {
      keyframes: {
        opacity: {
          "0%": { opacity: "0" },
          "100%": { opacity: "1" },
        },
      },
      animation: {
        opacity: "opacity 1s",
      },
    },
  },
  plugins: [],
};
