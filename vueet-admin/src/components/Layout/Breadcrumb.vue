<template>
    <div class="breadcrumb">
        <el-breadcrumb separator="/">
            <template v-for="(route, index) in breadcrumbs" :key="route.path">
                <el-breadcrumb-item v-if="index===breadcrumbs.length-1">{{route.meta.title}}</el-breadcrumb-item>
                <el-breadcrumb-item v-else="index===breadcrumbs.length-1" :to="{ path: route.path }">{{route.meta.title}}</el-breadcrumb-item>
            </template>
        </el-breadcrumb>
    </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, RouteLocationMatched } from 'vue-router'

/**
 * 设置面包屑
 * 1. 监听路由变化
 * 2. 通过 route.matched方法过滤
 */

const breadcrumbs = ref<RouteLocationMatched[]>([])

const route = useRoute()

const setBreadcrumbs = () => {
    breadcrumbs.value = route.matched.filter(
        (item) => item.meta && item.meta.title
    )
    // console.log("breadcrumbs: ",breadcrumbs.value)
}

watch(
    route,
    () => {
        setBreadcrumbs()
    },
    {
        immediate: true
    }
)

</script>

<style scoped lang="less">
.breadcrumb {
    margin-top: 48px;
    margin-bottom: 20px;
    margin-left: -24px;
    margin-right: -20px;
    padding: 24px;
    background-color: @white;
}

:deep(.el-breadcrumb__separator) {
    margin: 0 5px;
}
</style>