import { createI18n } from "vue-i18n"
import zhCn from "./zh-cn"
import zhTw from "./zh-cn"
import en from "./en"

const i18n = createI18n({
    legacy: false,
    locale: 'zhCn',
    messages: {
        zhCn,
        zhTw,
        en
    }
})

export default i18n