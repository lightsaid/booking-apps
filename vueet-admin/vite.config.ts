import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { fileURLToPath, URL } from "url";
import svgLoader from "vite-svg-loader";
// 按需导入 Element Plus 组件
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        svgLoader({ defaultImport: 'component' }),
        AutoImport({
            resolvers: [ElementPlusResolver()],
        }),
        Components({
            resolvers: [ElementPlusResolver()],
        }),
    ],
    resolve: {
        alias: {
            "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
    css: {
        preprocessorOptions: {
            less: {
                additionalData: "@import '@/assets/styles/index.less';",
            },
        },
    },
    server: {
        host: 'localhost',
        port: 3000,
        proxy: {
            '/api': {
                target: 'http://localhost:5000',
                rewrite: path => path.replace(/^\/api/, '')
            }
        }
    }
});
