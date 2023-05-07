import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue';
import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify';
import vuerouter from 'unplugin-vue-router/vite';
import layouts from 'vite-plugin-vue-layouts';
// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  build: { target: 'esnext', chunkSizeWarningLimit: 5000 },
  plugins: [
    vue({ template: { transformAssetUrls } }),
    vuerouter(),
    layouts(),
    vuetify({ autoImport: true }),
  ],
  css: { devSourcemap: true },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
