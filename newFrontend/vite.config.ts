import { fileURLToPath, URL } from "node:url";

import { defineConfig, loadEnv, splitVendorChunkPlugin } from "vite";
import { resolve, dirname } from "node:path";
import vue from "@vitejs/plugin-vue";
import vueI18n from "@intlify/vite-plugin-vue-i18n";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    splitVendorChunkPlugin(),
    vueI18n({
      runtimeOnly: false,
      // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
      compositionOnly: false,
      // you need to set i18n resource including paths !
      include: resolve(dirname(fileURLToPath(import.meta.url)), "./path/to/src/locales/**"),
    }),
  ],
  build:{
    reportCompressedSize:true,
    sourcemap: false,
    cssCodeSplit: true,
    chunkSizeWarningLimit: 1400,
    rollupOptions: {
      // https://rollupjs.org/guide/en/#outputmanualchunks
      output: {
        strict: true,
        manualChunks: {
          'general': [
          ],
          "data": [
          ],
        },
      }
    },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
      "~": fileURLToPath(new URL("./node_modules", import.meta.url)),
    },
  },
});
