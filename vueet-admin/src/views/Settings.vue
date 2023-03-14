<template>
    <div class="pageSettings">
        <el-tabs tab-position="left" :style="{height: offsetHeight + 'px'}" class="setting-tab">
            <el-tab-pane label="基本设置">
                <BasicSetting ref="basicRef" />
            </el-tab-pane>
            <el-tab-pane label="安全设置">
                <SecuritySetting />
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script setup lang="ts">
import BasicSetting from '@/components/BasicSetting.vue';
import SecuritySetting from '@/components/SecuritySetting.vue';
import { ref, watch, nextTick } from "vue"
const basicRef = ref<any>()
const offsetHeight = ref<number>(300)
watch(basicRef, ()=> {
    if (basicRef.value) {
        nextTick(()=>{
            offsetHeight.value = basicRef.value.$el.offsetHeight
        })
    }
}, {immediate: true})
</script>

<style scoped lang="less">
.pageSettings {
    padding: 20px 0;
}
</style>