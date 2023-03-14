<template>
    <aside class="vueet-layout-sidebar" ref="sidebarRef" :style="{ width: isCollapse ? '64px' : '210px' }">
        <section class="fixed-sidebar">
            <div class="menu">
                <div class="menu-box">
                    <!-- 开启路由模式 router -->
                    <el-menu :default-active="activeMenu" class="el-menu-vertical-demo" :collapse="isCollapse" router>
                       
                        <template v-for="route in routes" :key="route.path">
                            <!-- 存在2+子路由 -->
                            <el-sub-menu :index="route.path" v-if="route.children && route.children.length > 1 && !route?.meta?.hiddenMenu">
                                <template #title>
                                    <el-icon>
                                        <component :is="route?.meta?.icon" />
                                    </el-icon>
                                    <span>{{ route?.meta?.title }}</span>
                                </template>
                                <el-menu-item v-for="ch in route.children" :index="ch.path">{{ ch?.meta?.title }}</el-menu-item>
                            </el-sub-menu>

                            <!-- 有且仅有一个子路由 -->
                            <el-menu-item :index="route.children[0].path" v-else-if="route.children && route.children.length == 1 && !route?.meta?.hiddenMenu">
                                <el-icon> 
                                    <component :is="route.children[0]?.meta?.icon" /> 
                                </el-icon>
                                <template #title>{{ route.children[0]?.meta?.title }}</template>
                            </el-menu-item>

                            <!-- 没有子路由 -->
                            <el-menu-item v-else :index="route.path" v-if="!route?.meta?.hiddenMenu">
                                <el-icon>
                                    <component :is="route?.meta?.icon" />
                                </el-icon>
                                <template #title>{{ route?.meta?.title }}</template>
                            </el-menu-item>

                        </template>

                    </el-menu>
                </div>
            </div>

            <!-- 展开与收起 -->
            <div class="hamburger" @click="handleTrigger">
                <span>
                    <HamburgerIcon />
                </span>
            </div>
        </section>
    </aside>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router"
import { Menu as IconMenu, Location } from "@element-plus/icons-vue";
import HamburgerIcon from "@/assets/svgicon/hamburger.svg?component";
import { routes } from "@/router"
const sidebarRef = ref();
const isCollapse = ref(false);

console.log(routes)

const route = useRoute()
const activeMenu = computed(()=>{
    console.log(route.path)
    return route.path
})

const handleTrigger = () => {
    isCollapse.value = !isCollapse.value;
    console.log("width: ", sidebarRef.value.style.width);
};
</script>

<style scoped lang="less">
.vueet-layout-sidebar {
    transition: all 0.25s ease;
    .fixed-sidebar {
        position: fixed;
        display: flex;
        flex-direction: column;
        width: 210px;
        height: 100vh;
        max-height: 100vh;
        overflow-y: hidden;
        background-color: #f0f2f5;

        .menu {
            flex: 1;
            margin-top: 48px;
            overflow-y: auto;
            background: #fff;
            .scrollBar;

            .menu-box {
                height: 100%;
                .el-menu {
                    height: 100%;
                    border-right: none;
                }
            }
        }

        .hamburger {
            display: inline-flex;
            align-items: center;
            height: 40px;
            padding-left: 20px;
            border-top: solid 1px #dcdfe6;
            background: #fff;
            cursor: pointer;
            &:hover {
                background-color: #ecf5ff;
                span {
                    color: @primaryColor;
                }
            }
        }
    }

    el-menu-vertical-demo:not(.el-menu--collapse) {
        width: 100%;
        min-height: 100%;
    }
}
</style>
