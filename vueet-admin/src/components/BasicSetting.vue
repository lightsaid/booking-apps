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
                <el-upload class="avatar-uploader" name="files" :limit="1" accept="image/png, image/jpg, image/jpeg" 
                    multiple :action="uploadurl" :show-file-list="false"
                    :headers='{"Authorization": `Bearer ${token}`}'
                    :on-error="handleOnerror"
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
import { GetProfile, UpdateUser } from '@/api/user'
import { GetRoleById } from "@/api/other"
import { ProfileKey } from "@/store/profile"
import storage from "@/utils/storage"
const AdminCode = "ADMIN"

const uploadurl = import.meta.env.VITE_UPLOAD_URL
const imgBaseURL = import.meta.env.VITE_FILE_URL

const token = storage.get(ProfileKey)?.access_token
console.log(token)

const formSize = ref('default')
const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
    id: 0,
    name: '',
    phone_number: "",
    role: false,
    avatar: ""
})

const imageUrl = ref('')

onMounted(()=>{
    GetProfile().then(res=>{
        res.data.avatar = imgBaseURL + res.data.avatar
        ruleForm.id = res.data.id
        ruleForm.name = res.data.name
        ruleForm.phone_number = res.data.phone_number
        ruleForm.avatar =  res.data.avatar || ""
        imageUrl.value = res.data.avatar || ""
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
    let { data } = response
    imageUrl.value = data[0] || ruleForm.avatar
}

const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
    let allowTypes = ['image/jpeg', 'image/jpg', 'image/png']
    if (!allowTypes.includes(rawFile.type)) {
        ElMessage.error('Avatar picture must be JPG format!')
        return false
    } else if (rawFile.size / 1024 / 1024 > 2) {
        ElMessage.error('头像大小不能超过 2MB')
        return false
    }
    return true
}

const handleOnerror = (error: Error) => {
    try{
        let data = JSON.parse(error.message)
        ElMessage.error(data.msg)
    }catch(error) {
        ElMessage.error("上传失败")
    }
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
            const { id }  = ruleForm
            if (imageUrl.value[0] == ".") {
                imageUrl.value = imageUrl.value.substring(1)
            }
            UpdateUser(id, {...ruleForm, avatar: imageUrl.value} ).then(res=>{
                ElMessage.success(res.msg)
            })
        } else {
            console.log('error submit!', fields)
        }
    })
}

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