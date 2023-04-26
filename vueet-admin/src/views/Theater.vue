<template>
    <div class="page">
        <el-row :gutter="20" class="search-bar" justify="space-betwwen">
            <el-col :span="16">
                <el-form :inline="true" :model="searchForm" class="demo-form-inline">
                    <el-form-item label="电影院名">
                        <el-input v-model="searchForm.name" placeholder="" />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" @click="onSubmit">查询</el-button>
                    </el-form-item>
                </el-form>
            </el-col>
            <el-col :span="8" class="flex">
                <el-button type="primary" @click="newTheater">新增</el-button>
            </el-col>
        </el-row>

        <div class="wraper">
            <el-table :data="theaters.data" border style="width: 100%">
                <el-table-column prop="id" label="ID" width="120" />
                <el-table-column prop="name" label="电影院名" width="180" />
                <el-table-column prop="location" label="地理位置" />
                <el-table-column fixed="right" label="操作" width="120">
                <template #default="scope">
                    <el-button link type="primary" size="small" @click="seeHalls(scope.row)">
                        <!-- <router-link :to="`/theater/hell/${scope.row.id}`"></router-link> -->
                        查看大厅
                    </el-button>
                </template>
                </el-table-column>
            </el-table>
        </div>

        <div class="pagination">
            <el-pagination
                small
                background
                layout="prev, pager, next"
                :total="theaters.total"
                :page-size="theaters.page_size"
                model:current-page="theaters.page_num"
                class="mt-4"
            />
        </div>

        <div class="hall">
            <router-view></router-view>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref, watch } from "vue"
import { useTheater } from "@/composables/theaters"
import { useRouter } from "vue-router"

const router = useRouter()

const { theaters, getTheaters } = useTheater()

const searchForm = reactive({
    name: ""
})

let init_page: App.Pagination = {
    page_num: 1,
    page_size: 10
}

onMounted(()=> {
    getTheaters(init_page.page_num, init_page.page_size)
}) 

const seeHalls = (item: ApiRsp.TheaterResult) => {
    router.push(`/theater/hell/${item.id}`)
}

const newTheater = () => {
    alert("开发ing")
}

const onSubmit = () => {
    alert("开发ing")
}

</script>

<style scoped lang="less">
.page {
    min-height: 50vh;
    .search-bar {
        padding: 16px;
        .flex{
            display: flex;
            justify-content: end;
        }
    }
    .pagination{
        display: flex; 
        justify-content: center;
        margin-top: 20px;
    }
    .header{
        padding-bottom: 3px;
        border-bottom: 1px solid var(--el-color-primary);
    }
    .hall{
        margin-top: 10px;
    }
}
</style>