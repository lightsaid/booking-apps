<template>
    <div class="pageLogin">
        <div class="login-form">
            <div class="login-title">
                <MovieIcon width="32" height="32" />
                <p>Booking Apps</p>
            </div>

            <el-form ref="formRef" :model="numberValidateForm" label-width="100px">
                <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
                    <el-tab-pane label="手机号登录" name="first">
                        <el-form-item label="" label-width="0" prop="age" :rules="[
                            { required: true, message: 'age is required' },
                            { type: 'number', message: 'age must be a number' },
                        ]">
                            <el-input v-model.number="numberValidateForm.age" type="text" placeholder="请输入手机号"
                                autocomplete="off" />
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
                            <el-button @click="resetForm(formRef)">Reset</el-button>
                        </el-form-item>
                    </el-tab-pane>
                    <el-tab-pane label="账号密码登录" name="second">
                        <el-form-item label="" label-width="0" prop="age" :rules="[
                            { required: true, message: 'age is required' },
                            { type: 'number', message: 'age must be a number' },
                        ]">
                            <el-input v-model.number="numberValidateForm.age" type="text" placeholder="请输入手机号"
                                autocomplete="off" />
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="submitForm(formRef)">Submit</el-button>
                            <el-button @click="resetForm(formRef)">Reset</el-button>
                        </el-form-item>
                    </el-tab-pane>
                </el-tabs>

            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import type { TabsPaneContext, FormInstance } from "element-plus"
import MovieIcon from "@/assets/movie.svg"

const activeName = ref('first')
const formRef = ref<FormInstance>()
const handleClick = (tab: TabsPaneContext, event: Event) => {
    console.log(tab, event)
}

const numberValidateForm = reactive({
    age: '',
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>

<style scoped lang="less">
.pageLogin {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: url(../assets/loginbg.jpg) no-repeat;
    background-size: cover;
    overflow: hidden;


    .login-form {
        min-width: 320px;
        margin-top: -100px;
        padding: 20px;
        border-radius: 6px;
        background: @white;

        .login-title {
            display: flex;
            justify-content: center;
            align-items: center;
            margin-bottom: 2em;

            p {
                margin-left: 10px;
                font-size: 1.5rem;
                font-weight: 500;
            }
        }

        /deep/.el-tabs__nav-scroll {
            overflow: hidden;
            display: flex;
            justify-content: center;
        }
    }
}</style>