import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import { ElMessage } from "element-plus"
import 'element-plus/dist/index.css'
import i18n from "./language"
import { createPinia } from 'pinia'
import "./assets/styles/base.less";


const app = createApp(App);
app.use(router);
app.use(i18n)
app.use(createPinia())

app.config.globalProperties.$message = ElMessage

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component);
}


app.mount("#app");
