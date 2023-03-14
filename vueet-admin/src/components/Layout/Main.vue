<template>
    <!-- App Main 主体 -->
    <session class="vueet-layout-main">
        <!-- 顶部 Header Fixed -->
        <Navbar />

        <!-- Main 路由入口 -->
        <section class="">
            <!-- Breadcrumb 面包屑 -->
            <div class="page">
                <Breadcrumb />
                <div class="app-main">
                    <router-view v-slot="{ Component, route }">
                        <transition name="fade-transform" mode="out-in">
                            <keep-alive>
                                <component :is="Component" :key="route.path" />
                            </keep-alive>
                        </transition>
                    </router-view>
                </div>
                <Footer />
            </div>

        </section>
    </session>
</template>

<script setup lang="ts">
import Navbar from "./Navbar.vue";
import Footer from "./Footer.vue"
import Breadcrumb from "./Breadcrumb.vue";
</script>

<style scoped lang="less">
.vueet-layout-main {
    position: relative;
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    overflow-y: auto;
    background-color: #f0f2f5;
    border-left: 1px solid #dcdfe6;
    .page {
        position: relative;
        padding: 0 24px 24px 24px;
    }
    .app-main {
        border-radius: 6px;
        background: white;
        .bxShadow;
    }
}
</style>
