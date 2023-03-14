<template>
    <div>
        <div v-for="i in 1">
            <p> {{i}} Dashboard Page </p>
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
