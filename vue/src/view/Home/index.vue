<template>

  <div>
    <a href="https://vitejs.dev" target="_blank">
      <!-- /vite.svg 会默认去public下找 public下的东西好像不会被打包 -->
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <!-- img的src自动可以用自己配的alias https://github.com/vitejs/vite/tree/main/packages/plugin-vue -->
      <img src="@/assets/logo/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>

  <div class="text-left">
    <div class="bg-yellow-300" v-for="item in bloglist" :key="item.ID">
      <div>用户:{{item.user_id}}</div>
      <div>创建时间:{{item.CreatedAt}}</div>
      <div v-html="item.Content"></div>
    </div>
  </div>

</template>

<script lang='ts' setup name="home">
// {
// "ID": 2,
// "CreatedAt": "2022-09-19T16:23:26.081+08:00",
// "UpdatedAt": "2022-09-19T16:23:26.081+08:00",
// "DeletedAt": null,
// "user_id": "17",
// "Content": "弹道内膜癌烦恼烦恼岸防泼",
// "Tag": ""
// }
const bloglist = ref<any>([])
//获取列表
async function getBlogList() {
  const { data } = await http.get("/editorAllList")
  bloglist.value = data.data
}
getBlogList()

</script> 

<style lang='less' scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #008d47aa);
}

.hello_word_box {
  background-color: red;
}
</style>