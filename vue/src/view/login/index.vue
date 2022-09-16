<template>
<div class="w-100vw h-100vh content">

    <div class="box" w-200>
      <el-card class="w-100% h-100%" shadow="always" :body-style="{ padding: '20px' }">
        <template #header>
          <div>
            <span>登录</span>
          </div>
        </template>
        <div>
          <el-form ref="ruleFormRef" :model="userInfo" :rules="rules" label-width="80px" status-icon>
    
            <el-form-item label="帐号" prop="userName">
              <el-input v-model="userInfo.userName" />
            </el-form-item>
            <el-form-item label="验证码" prop="code">
              <el-input v-model="userInfo.code" />
            </el-form-item>
            <el-form-item label="密码" prop="passWord">
              <el-input v-model="userInfo.passWord" />
            </el-form-item>
    
            <el-form-item>
              <el-button type="primary" @click="submitForm(ruleFormRef)">登录</el-button>
              <el-button @click="resetForm(ruleFormRef)">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-card>
    </div>

</div>

</template>

<script lang='ts' setup name="login">
import type { FormInstance, FormRules } from 'element-plus'


const ruleFormRef = ref<FormInstance>()
const userInfo = reactive<{
  userName: string;
  code: string;
  passWord: string;
}>({
  userName: "",
  code: "",
  passWord: "",
})
const rules = reactive<FormRules>({
  userName: [
    { required: true, message: '必填项', trigger: 'blur' },
    { min: 3, max: 15, message: '长度在 3 到 15 之间', trigger: 'blur' },
  ],
  code: [
    { required: true, message: '必填项', trigger: 'change', },
  ],
  passWord: [
    { required: true, message: '必填项', trigger: 'change', },
    { min: 3, max: 15, message: '长度在 3 到 15 之间', trigger: 'blur' },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      // http.post("/captcha",{},{
      //   headers:{
      //     loading: true,
      //     nprogress: true,
      //   }
      // })
      console.log(123)
      router.push("/layout")
    } else {
      // console.log('error submit!', fields)
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

</script> 

<style lang='less' scoped>
@import "./index.less";
</style>