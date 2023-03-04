<template>
    <div>
        <div v-for="i in 100">
            <p> {{i}} Dashboard Page <span>Need Auth Required</span></p>
        </div>
    </div>
</template>

<script setup lang="ts"></script>

<style scoped lang="less">
div {
    padding: 12px;
}
span {
    color: @primaryColor;
}
</style>
