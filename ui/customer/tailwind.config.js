/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{js,jsx,ts,tsx}"],
    theme: {
      extend: {
        colors: {
          prim: {
            50: "#E7F5FF",
            100: "#D0EBFF",
            200: "#A5D8FF",
            300: "#74C0FC",
            400: "#4DABF7",
            500: "#339AF0",
            600: "#228BE6",
            700: "#1C7ED6",
            800: "#1971C2",
            900: "#1864AB"
          }
        }
      }
    },
    plugins: []
  };