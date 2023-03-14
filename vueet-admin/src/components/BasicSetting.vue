<template>
    <div class="basicSetting">
        <p class="title">基本设置</p>
        <el-form ref="ruleFormRef" :model="ruleForm" :rules="rules" label-width="120px" :size="formSize" status-icon>
            <el-form-item label="用户名" prop="name">
                <el-input v-model="ruleForm.name" />
            </el-form-item>
            <el-form-item label="手机号" prop="phone_number">
                <el-input disabled v-model="ruleForm.phone_number" />
            </el-form-item>
            <el-form-item label="管理员" prop="role">
                <el-switch v-model="ruleForm.role" />
            </el-form-item>
            <el-form-item label="头像" prop="avatar">
                <el-upload class="avatar-uploader" accept="image/png, image/jpg, image/jpeg" action="/upload" :show-file-list="false"
                    :on-success="handleAvatarSuccess" :before-upload="beforeAvatarUpload">
                    <img v-if="ruleForm.avatar" :src="ruleForm.avatar" class="avatar" />
                    <el-icon v-else class="avatar-uploader-icon">
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" class="submit" @click="submitForm(ruleFormRef)">保存</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { ElMessage } from 'element-plus'
import { Plus } from '@element-plus/icons-vue'
import type { UploadProps } from 'element-plus'
import { GetProfile } from '@/api/user'
import { GetRoleById } from "@/api/other"

const AdminCode = "ADMIN"

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
    name: '',
    phone_number: "",
    role: false,
    avatar: ""
})

const imageUrl = ref('')

onMounted(()=>{
    GetProfile().then(res=>{
        ruleForm.name = res.data.name
        ruleForm.phone_number = res.data.phone_number
        ruleForm.avatar = res.data.avatar || ""
        // 查询是否是管理员
        GetRoleById(res.data.role_id).then(res=>{
            ruleForm.role = res.data.code === AdminCode
        })
    })
})

const handleAvatarSuccess: UploadProps['onSuccess'] = (
    response,
    uploadFile
) => {
    ruleForm.avatar = URL.createObjectURL(uploadFile.raw!)
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
    if (rawFile.type !== 'image/jpeg') {
        ElMessage.error('Avatar picture must be JPG format!')
        return false
    } else if (rawFile.size / 1024 / 1024 > 2) {
        ElMessage.error('Avatar picture size can not exceed 2MB!')
        return false
    }
    return true
}

const rules = reactive<FormRules>({
    name: [
        { required: true, message: 'Please input Activity name', trigger: 'blur' },
        { min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
    ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
    if (!formEl) return
    await formEl.validate((valid, fields) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!', fields)
        }
    })
}

const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
}

const options = Array.from({ length: 10000 }).map((_, idx) => ({
    value: `${idx + 1}`,
    label: `${idx + 1}`,
}))
</script>

<style scoped lang="less">
.basicSetting {
    padding: 0 24px;

    .title {
        font-size: 20px;
    }

    .el-form {
        width: 40%;
        margin-top: 30px;
    }

    .avatar-uploader .avatar {
        width: 128px;
        height: 128px;
        display: block;
    }

    .avatar-uploader .el-upload {
        border: 1px dashed var(--el-border-color);
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
        transition: var(--el-transition-duration-fast);
    }

    .avatar-uploader .el-upload:hover {
        border-color: var(--el-color-primary);
    }

    .el-icon.avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 128px;
        height: 128px;
        text-align: center;
        border-radius: 6px;
        .bxShadow;
    }
}
</style>