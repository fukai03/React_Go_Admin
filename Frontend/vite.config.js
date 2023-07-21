/* eslint-disable no-undef */
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            '/api': {
                target: 'http://localhost:8888',
                changeOrigin: true,
            },
            '/ai': {
                target: 'http://10.11.135.27:8100',
                rewrite: (path) => path.replace(/^\/ai/, ''),
                changeOrigin: true,
            }
        },
        hmr: true
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src'),
            'utils': path.resolve(__dirname, './src/utils'),
        }
    }
})
