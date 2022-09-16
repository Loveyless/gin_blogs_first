<template>
  <div class="m-4 box">
    <div>
      我是子元素
    </div>

    <div> 父元素的默认插槽都在我下面👇</div>
    <slot name="default" title="123" msg="默认插槽的传值">默认内容 但是父组件传插槽了 这个默认内容不会显示</slot> <!-- 作用域插槽也是插槽传值 除了name都可以拿到 -->

    <div>具名在这里👇</div>
    <slot name="IHaveName" msg="Name插槽的传值" v-bind="obj">Name默认内容</slot> <!-- 作用域插槽也是插槽传值 除了name都可以拿到 -->


    <div>或者在这里👇</div>
    <slot name="IHaveMony" msg="Mony插槽的传值">Mony默认内容</slot> <!-- 作用域插槽也是插槽传值 除了name都可以拿到 -->

  </div>
</template>

<script lang='ts' setup>

//作用域插槽之间用就是直接映射key value
const obj = ref<any>({
  age: 19,
  sex: "male"
})
</script> 

<style lang='less' scoped>
.box {
  background-color: aquamarine;
}
</style>