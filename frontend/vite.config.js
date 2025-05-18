import { defineConfig } from 'vite'
import vuePlugin from "@vitejs/plugin-vue";

export default defineConfig({
 resolve: {
  alias: {
   '@': '/src',
  },
 },
 plugins: [
  vuePlugin({
   template: {
    transformAssetUrls: {
     base: null,
     includeAbsolute: false,
    },
   },
  }),
 ],
})