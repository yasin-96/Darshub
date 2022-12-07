import { fileURLToPath, URL } from "node:url";

import { defineConfig, loadEnv, splitVendorChunkPlugin } from "vite";
import { resolve, dirname } from "node:path";
import vue from "@vitejs/plugin-vue";
import vueI18n from "@intlify/vite-plugin-vue-i18n";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 5173,
  },
  plugins: [
    vue(),
    splitVendorChunkPlugin(),
    vueI18n({
      runtimeOnly: false,
      // to use Vue I18n Legacy API
      compositionOnly: false,
      // you need to set i18n resource including paths !
      include: resolve(dirname(fileURLToPath(import.meta.url)), "./path/to/src/locales/**"),
    }),
  ],
  css: {
    postcss: {
      plugins: [
        require("tailwindcss/nesting"),
        require("tailwindcss")({
          config: process.env?.TAILWIND_CONFIG ? `tailwind-${process.env.TAILWIND_CONFIG}.config.js` : "./tailwind.config.js",
        }),
        require("autoprefixer"),
      ],
    },
  },
  build: {
    reportCompressedSize: true,
    sourcemap: false,
    cssCodeSplit: true,
    chunkSizeWarningLimit: 1400,
    // rollupOptions: {
    //   // https://rollupjs.org/guide/en/#outputmanualchunks
    //   output: {
    //     strict: true,
    //     manualChunks: {
    //       general: [],
    //       data: [],
    //     },
    //   },
    // },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "~": fileURLToPath(new URL("./node_modules", import.meta.url)),
    },
  },
});

// chainWebpack: (config) => {
//   config
//   .plugin('define')
//   .tap((args) => {
//     const gitRevisionPlugin = new GitRevisionPlugin({
//       commithashCommand: 'rev-list --max-count=1 --no-merges HEAD',
//       versionCommand: 'describe --always --tags',
//     });

//     args[0]['process.env']['VUE_APP_GIT_VERSION'] = JSON.stringify(gitRevisionPlugin.version());
//     args[0]['process.env']['VUE_APP_GIT_HASH'] = JSON.stringify(gitRevisionPlugin.commithash());
//     args[0]['process.env']['VUE_APP_GIT_LAST_COMMIT_DATE'] = JSON.stringify(gitRevisionPlugin.lastcommitdatetime());
//     return args;
//   });

//   config.module.rule('pdf')
//     .test(/\.pdf$/)
//     .use('file-loader').loader('file-loader');
// },
