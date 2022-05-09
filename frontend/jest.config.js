module.exports = {
  preset: "@vue/cli-plugin-unit-jest",
  snapshotSerializers: ["<rootDir>/node_modules/jest-serializer-vue"],
  moduleFileExtensions: ["js", "vue"],
  moduleNameMapper: {
    "^@/(.*)$": "<rootDir>/src/$1",
  },
  transform: {
    // process `*.vue` files with `vue-jest`
    ".*\\.(js)$": "babel-jest",
    ".*\\.(vue)$": "vue-jest",
  },
  collectCoverage: true,
  collectCoverageFrom: ["**/*.{js,vue}", "!**/node_modules/**"],
};
