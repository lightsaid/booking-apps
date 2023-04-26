<template>
    <div class="page">
        <header class="header">
            <span style="color:var(--el-color-primary); font-weight: bold;">{{ theater?.name }}</span>
            <span style="display: inline-block; padding-left: 4px; color: var(--el-color-primary);">大厅</span>
        </header>
        <div class="wraper">
            <el-table :data="halls.data" style="width: 100%">
                <el-table-column prop="id" label="ID" width="120" />
                <el-table-column prop="name" label="大厅名" />
                <el-table-column fixed="right" label="操作" width="120">
                </el-table-column>
            </el-table>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue"
import { useRoute } from "vue-router"
import { useHall, useTheater } from "@/composables/theaters"

const route = useRoute()

const { halls, getHalls } = useHall()
const { theater, getTheater }  = useTheater()

const init_page = { page_num: 1, page_size: 100 }

onMounted(() => {
    if (route.params.id) {
        getHalls(Number(route.params.id), init_page.page_num, init_page.page_size)
        getTheater(Number(route.params.id))
    }
})


</script>

<style scoped lang="less">
.page {
    min-height: 50vh;

    .wraper {
        margin-top: 10px;
    }
}
</style>