module.exports = {
    tabWidth: 2,
    useTabs: false,
    trailingComma: "none",
    semi: true,
    singleQuote: false,
    plugins: [require("prettier-plugin-tailwindcss")],
    tailwindConfig: "./tailwind.config.js"
};