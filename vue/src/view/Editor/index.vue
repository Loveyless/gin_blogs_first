<template>
  <div class="content">
    <WangEditor ref="editorRef" height="700px"></WangEditor>
    <div class="box">
      <el-button type="primary" plain @click="showHtml">预览</el-button>
      <el-button type="primary" plain @click="submitForm">发布</el-button>
    </div>
  </div>
</template>

<script lang='ts' setup name="wangeditor">
// import WangEditor from "@/components/WangEditor/index.vue"
const global = GlobalStore()

//编辑器实例
const editorRef = ref<any>('')

//预览
function showHtml() {

}

onMounted(() => {
  console.log(1)
})

//发布
async function submitForm() {
  //发送请求
  const { data } = await http.post('/editor', {
    user_id: global.Id,
    content: editorRef.value.valueHtml,
  })
  if (data.status == 200) {
    editorRef.value.setEditorHtml("")
  }
}

onMounted(() => {
  console.log(9)
})
</script> 

<style lang="less" scoped>
@import "./index.less";
</style>