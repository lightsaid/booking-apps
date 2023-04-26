<template>
    <div class="pageLogin">
        <div class="login-form">
            <div class="login-title">
                <MovieIcon width="32" height="32" />
                <p>Booking Apps</p>
            </div>

            <el-form ref="formRef" :model="loginForm" label-width="0">
                <el-tabs v-model="loginForm.login_type" @tab-click="handleClick">
                    <el-tab-pane label="短信验证码登录" name="CODE">
                        <template v-if="loginForm.login_type === 'CODE'">
                            <el-form-item label="" prop="phone_number" :rules="[
                                { validator: checkPhone, trigger: 'blur' },
                            ]">
                                <el-icon>
                                    <Iphone />
                                </el-icon>
                                <el-input v-model.trim="loginForm.phone_number" type="text" placeholder="请输入手机号"
                                    autocomplete="off" />
                            </el-form-item>

                            <el-row :gutter="20">
                                <el-col :span="14">
                                    <el-form-item label="" prop="code" :rules="[
                                        { required: true, message: '请输入短信验证码' },
                                    ]">
                                        <el-icon>
                                            <Lock />
                                        </el-icon>
                                        <el-input v-model.number="loginForm.code" type="text" placeholder="请输验证码"
                                            autocomplete="off" />
                                    </el-form-item>
                                </el-col>
                                <el-col :span="10">
                                    <el-button :disabled="smsStatus" type="primary"
                                        @click="handleSendSms">{{ smsText }}</el-button>
                                </el-col>
                            </el-row>

                            <el-form-item>
                                <el-button type="primary" class="submit" @click="submitForm(formRef)">登录</el-button>
                            </el-form-item>
                        </template>
                    </el-tab-pane>
                    <el-tab-pane label="账号密码登录" name="PASS">
                        <template v-if="loginForm.login_type === 'PASS'">
                            <el-form-item label="" prop="phone_number" :rules="[
                                { validator: checkPhone, trigger: 'blur' },
                            ]">
                                <el-icon>
                                    <Iphone />
                                </el-icon>
                                <el-input v-model.trim="loginForm.phone_number" type="text" placeholder="请输入手机号"
                                    autocomplete="off" />
                            </el-form-item>

                            <el-form-item label="" prop="password" :rules="[
                                { required: true, message: '请输入密码' },
                                { min: 6, max: 16, message: `长度必须是 6～16 个字符`, trigger: 'blur' },
                            ]">
                                <el-icon>
                                    <Lock />
                                </el-icon>
                                <el-input v-model="loginForm.password" type="password" placeholder="请输入密码"
                                    autocomplete="off" />
                            </el-form-item>
                            <el-form-item>
                                <el-button type="primary" class="submit" @click="submitForm(formRef)">登录</el-button>
                            </el-form-item>
                        </template>
                    </el-tab-pane>
                </el-tabs>
            </el-form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { TabsPaneContext, FormInstance } from "element-plus"
import MovieIcon from "@/assets/svgIcon/movie.svg?component"
import { Login } from "../api/user"
import { SendSMS } from "../api/other"
import { useProfileStore } from "@/store"
const smsText = ref('获取验证码')
const smsStatus = ref(true)
const formRef = ref<FormInstance>()

const store = useProfileStore()
const router = useRouter()

const handleClick = (tab: TabsPaneContext, event: Event) => { }

const loginForm = reactive({
    phone_number: localStorage.getItem("login_phone_number") || '',
    code: '',
    password: '',
    login_type: 'CODE',
})

watch(
    loginForm,
    ()=>{
        var re = /^1\d{10}$/
        if (re.test(loginForm.phone_number)) {
           smsStatus.value = false
        }else{
            smsStatus.value = true
        }
    },
    {
        immediate: true
    }
)

const checkPhone = (rule: any, value: any, callback: any) => {
    if (value === "") {
        return callback(new Error('请输入手机号码'))
    } else {
        var re = /^1\d{10}$/
        if (!re.test(value)) {
            return callback(new Error('请输入11位手机号'))
        }
    }
    callback()
}

const handleSendSms = () => {
    let count = 60
    smsText.value = `${count}秒后重新获取`
    smsStatus.value = true
    let cur: number
    let timer = setInterval(() => {
        cur = count--
        if (cur <= 0) {
            cur = 0
        }
        smsText.value = `${cur}秒后重新获取`
        if (count <= 0) {
            smsStatus.value = false
            smsText.value = "获取验证码"
            clearInterval(timer)
        }
    }, 1000)

    SendSMS({phone_number: loginForm.phone_number})
}

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            localStorage.setItem("login_phone_number", loginForm.phone_number);
            if (loginForm.login_type === "PASS"){
                Reflect.deleteProperty(loginForm, "code")
            }
            Login(loginForm).then(res=>{
                store.setProfile(res.data)
                router.replace({path:"/"})
            })
        } else {
            console.log('error submit!')
            return false
        }
    })
}

</script>

<style scoped lang="less">
.pageLogin {
    width: 100vw;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: url(../assets/imgs/loginbg.jpg) no-repeat;
    background-size: cover;
    overflow: hidden;

    .login-form {
        width: 380px;
        margin-top: -100px;
        padding: 20px;
        border-radius: 6px;
        background: @white;

        .el-form {
            padding: 0 15px
        }

        .login-title {
            display: flex;
            justify-content: center;
            align-items: center;
            margin: 1em 0 1.2em 0;

            p {
                margin-left: 10px;
                font-size: 1.5rem;
                font-weight: 500;
                color: @text;
            }
        }

        .el-button {
            width: 100%;
        }

        .submit {
            font-size: 1rem;
        }
    }

    :deep(.el-tabs__content) {
        margin-top: 1.5em;
    }

    :deep(.el-tabs__nav-scroll) {
        overflow: hidden;
        display: flex;
        justify-content: center;

        .el-tabs__item {
            font-weight: normal;

            .is-active {
                color: var(--el-color-primary);
            }
        }
    }

    :deep(.el-form-item) {
        position: relative;
        margin-bottom: 1.8rem;

        .el-input__inner {
            padding-left: 20px;
        }

        .el-icon {
            height: 100%;
            position: absolute;
            left: 10px;
            top: 0;
            z-index: 3;
        }
    }
}</style>