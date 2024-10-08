import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import fs from 'fs'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  build: {
    outDir: '../dist/files',
    emptyOutDir: true,
  },
  server: {
    https: {
      key: fs.readFileSync(path.resolve(__dirname, 'vite.key')),
      cert: fs.readFileSync(path.resolve(__dirname, 'vite.crt')),
    },
    port: 3000,
  },
})
