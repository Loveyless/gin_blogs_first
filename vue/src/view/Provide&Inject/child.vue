<template>
  我是后代组件 注入的数据

<br>

  普通接收 count->{{count1}}
  <button @click="count1++">+1</button>

<br>

  ref接收 count->{{count2}}
  <button @click="count2++">+1</button>

<br>

  {{info}}

<br>

  响应式注入
  {{reactivity}}

  <br>
  --------------------------
  <br>
  <img src="https://cdn.jsdelivr.net/gh/Loveyless/img-clouding/img/20220913224831.png" />

</template>

<script lang='ts' setup>

const count1: any = inject("count")
const count2: any = ref(inject("count"))
const info: any = inject("info")
const reactivity: any = inject("reactivity")
</script> 
