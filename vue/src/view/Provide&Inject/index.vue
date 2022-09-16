<template>
提供的数据
{{count}}
<button @click="count++">+1</button>

<br>

{{info}}

<br>

响应式注入
{{reactivity}}
<button @click="reactivity++">修改响应式注入</button>

<br>
--------------------------
<br>

<Child></Child>

</template>

<script lang='ts' setup name="provide&inject">

const count = ref<number>(0)
const info = reactive({
  name:"张三",
  age:19
}) 
let reactivity = ref<number>(999)
provide("count", count.value)
provide("info", info)
provide("reactivity", computed(() => reactivity))

</script>