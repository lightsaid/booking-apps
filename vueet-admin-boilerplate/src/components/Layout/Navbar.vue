<template>
    <div>
        <section class="vueet-layout-navbar">
            <div class="logo">
                <span>
                    <Logo width="32px" height="32px" />
                </span>
                <span class="logo-text">Booking&nbsp;Apps</span>
            </div>
            <ul class="navbar">
                <li class="notice">
                    <el-badge :value="3" class="item">
                        <el-icon>
                            <Bell />
                        </el-icon>
                    </el-badge>
                </li>
                <li class="box-settings">
                    <div class="navbar-item">
                        <el-avatar :src="Avatar" :size="28" />
                        <span class="name">LightSaid</span>
                    </div>
                    <div class="navbar-options">
                        <div>
                            <el-icon>
                                <Setting />
                            </el-icon>
                            <span class="text">{{ t("setting") }}</span>
                        </div>
                        <div>
                            <el-icon>
                                <SwitchButton />
                            </el-icon>
                            <span class="text">{{ t("logout") }}</span>
                        </div>
                    </div>
                </li>
                <li class="box-settings language" @mouseover="handleMouseover">
                    <div class="navbar-item">
                        <Translation />
                    </div>
                    <div class="navbar-options" v-show="hovering">
                        <div @click="changLang(en)">
                            <el-icon>
                                🇺🇸 
                            </el-icon>
                            <span class="text">English</span>
                        </div>
                        <div @click="changLang(zhCn)">
                            <el-icon>
                                🇨🇳
                            </el-icon>
                            <span class="text">简体中文</span>
                        </div>
                        <div @click="changLang(zhTw)">
                            <el-icon>
                                🇭🇰
                            </el-icon>
                            <span class="text">繁體中文</span>
                        </div>
                    </div>
                </li>
            </ul>
        </section>
    </div>
</template>

<script setup lang="ts">
import Logo from "@/assets/svgicon/movie.svg?component"
import Translation from "@/assets/svgicon/translate.svg?component"
import Avatar from "@/assets/imgs/xxx.png"
import { ref } from "vue"
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import zhTw from 'element-plus/dist/locale/zh-tw.mjs'
import en from 'element-plus/dist/locale/en.mjs'
import { useI18n } from "vue-i18n"
const { t } = useI18n()
const {locale: i18nLocal} = useI18n()

const locale = ref(zhCn)
const hovering = ref(false)

const handleMouseover = () => {
    hovering.value = true
}

const changLang = (language: any) => {
	locale.value = language
	i18nLocal.value = language.name
    hovering.value = false
}
</script>

<style scoped lang="less">
.vueet-layout-navbar {
    display: flex;
    align-items: center;
    width: 100%;
    height: 48px;
    position: fixed;
    top: 0;
    right: 0;
    z-index: 9;
    padding: 0 16px;
    background-color: rgba(0, 0, 0, 1);

    .logo {
        display: flex;
        align-items: center;
        color: white;
        .logo-text {
            padding-left: 12px;
            font-size: 22px;
            font-weight: 500;
        }
    }

    .navbar {
        display: flex;
        height: 100%;
        margin-left: auto;
        margin-right: 20px;
        color: white;
        li {
            position: relative;
            display: inline-flex;
            align-items: center;
            height: 100%;
            padding: 0 10px;
            cursor: pointer;
            &:hover {
                background-color: rgba(255, 255, 255, .2);
            }
            .navbar-item {
                display: flex;
                justify-content: center;
                align-items: center;
                font-size: 15px;
                .name {
                    display: inline-block;
                    margin-left: 7px;
                }
            }
            &.box-settings:hover{
                .navbar-options{
                    display: block;
                    animation: fade-in .2s both;
                }
            }
            .navbar-options {
                display: none;
                position: absolute;
                bottom: 0;
                left: 1px;
                font-size: 14px;
                color: #666;
                background-color: #fff;
                border-radius: 4px;
                animation: fade-out 0.2s both;
                .bxShadow;
                div {
                    display: flex;
                    align-items: center;
                    min-width: 150px;
                    margin-top: 10px;
                    padding: 0 12px;
                    height: 30px;
                    line-height: 30px;
                    &:last-child{
                        margin-bottom: 10px;
                    }
                    &:hover{
                        background-color: rgba(0, 0, 0, .1);
                    }
                    .el-icon {
                        font-size: 16px;
                    }
                    .text {
                        padding-left: 10px;
                    }
                }
            }
        }
        .notice {
            padding: 0 20px;
            padding-top: 5px;
        }
        .language{
            .navbar-item {
                font-size: 18px;
                padding: 0 10px;
            }
            .navbar-options{
                left: -105%;
                justify-content: center;
                transform: translateX(-100%);
                div{
                    min-width: 120px;
                    justify-content: center;
                }
            }
        }
    }
}


</style>
