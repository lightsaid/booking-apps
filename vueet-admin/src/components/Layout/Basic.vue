<template>
    <div class="vueet-layout">
        <section class="vueet-layout-container">
            <!-- 侧边栏 -->
            <Sidebar :isCollapse="isCollapse" />

            <!-- 主体 -->
            <Main />
        </section>
    </div>
</template>

<script setup lang="ts">
import Sidebar from "./Sidebar.vue";
import Main from "./Main.vue";
import { ref } from "vue";

const isCollapse = ref(true);
</script>

<style scoped lang="less">
.vueet-layout {
    .vueet-layout-container {
        display: grid;
        grid-template-columns: auto 1fr;
    }
}
</style>
