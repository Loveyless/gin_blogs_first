<template>
  <div class="w-100vw h-100vh content">

    <div class="box" w-200>
      <el-card class="w-100% h-100%" shadow="always" :body-style="{ padding: '20px' }">
        <template #header>
          <div>
            <div>博客管理系统</div>
            <div>注册</div>
          </div>
        </template>
        <div>
          <el-form ref="ruleFormRef" :model="userInfo" :rules="rules" label-width="50px" status-icon>

            <el-form-item label="帐号" prop="username">
              <el-input v-model="userInfo.username">
                <template #prefix>
                  <el-icon>
                    <i-carbon-user />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input v-model="userInfo.password">
                <template #prefix>
                  <el-icon>
                    <i-carbon-password />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="submitForm(ruleFormRef)" plain>注册</el-button>
              <el-button @click="resetForm(ruleFormRef)">重置</el-button>
              <el-button @click="() => ($router.push('/login'))">返回登录页</el-button>
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
  username: string;
  password: string;
}>({
  username: "123",
  password: "123",
})
const rules = reactive<FormRules>({
  username: [
    { required: true, message: '必填项', trigger: 'blur' },
    { min: 3, max: 15, message: '长度在 3 到 15 之间', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '必填项', trigger: 'change', },
    { min: 3, max: 15, message: '长度在 3 到 15 之间', trigger: 'blur' },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      const data = await http.post("/user", userInfo, {
        headers: {
          loading: false,
          nprogress: false,
        }
      })

      // router.push("/layout")
    } else {
      console.log('error submit!', fields)
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