<template>
  <!-- {{ $router.options.routes[1].children }} -->
  <el-menu default-active="/layout/home" mode="horizontal" menu-trigger="click" router>
    <el-sub-menu index="">
      <template #title>vue api</template>
      <el-menu-item v-for="item in $router.options.routes[2].children" :index="'/layout/' + item.path">{{ item.name }}
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script lang='ts' setup>

// const router = useRouter()
// 拿到路由表
// const routes = ref(router.options.routes[1].children)
// console.log(routes.value);

</script> 

<style lang='less' scoped>
</style>