const colors = require("tailwindcss/colors");
const { Icons } = require("tailwindcss-plugin-icons");


/**
 * @type {import('tailwindcss-plugin-icons').Options}
 */
const iconsOptions = ({ theme }) => ({
  heroicons: {
    icons: {},
    scale: 1.5,
    includeAll: true,
    location: "https://esm.sh/@iconify-json/heroicons@1.1.6/icons.json",
  },
});

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx,html}"],
  theme: {
    //https://tailwindcss.com/docs/theme#extending-the-default-theme
    extend: {},
    colors: { ...colors },
  },
  plugins: [
    require("@tailwindcss/typography"),
    require("@tailwindcss/forms"),
    require("@tailwindcss/line-clamp"),
    require("@tailwindcss/aspect-ratio"),
    Icons(iconsOptions),
    require("daisyui")
  ],
};
