<template>
  <div class="father_box">

    <div>官网插槽有个例子还不错可以看看在最下面</div>
    <div>插槽可以是任意合法的模板内容 !!甚至是组件!!</div>
    <div>插槽内容可以访问到父组件的数据作用域，因为插槽内容本身是在父组件模板中定义的。</div>
    <div>比如这个msg和下面插槽里的{{ msg }} </div>
    <div>插槽内容无法访问子组件的数据，请牢记一条规则：</div>

    <div>我是父元素</div>

    <!-- v-slot 有对应的简写 #，因此 <template v-slot:header> 可以简写为 <template #header></template> -->
    <!-- 没有提供 name 的 <slot> 插口会隐式地命名为“default”。 -->
    <SonSlot>
      <template #default="props">
        我是父组件的插槽 我插进来了 我是默认插槽 {{ msg }}{{ props }}
        <!-- 作用域插槽拿到传的值 -->
      </template>

      <template #[SlotName]="props">
        <!-- 具名插槽 + 动态插槽 -->
        我是父组件的插槽 我插进来了 我是具名动态插槽 {{ msg }}{{ props }}
        <!-- 作用域插槽拿到传的值 -->
      </template>
    </SonSlot>


    <div>
      <button @click="changeSlot">动态插槽</button>
    </div>


    <!-- 只封装了方法的插槽 -->
    <MouseTracker #default="{ x, y }">
      捕获当前鼠标位置 pageX{{ x }}|| pageY{{ y }}
    </MouseTracker>


  </div>
</template>

<script lang='ts' setup>

const msg = ref<string>("父组件的hello word")

//动态插槽名
const SlotName = ref<string>("IHaveName")
function changeSlot() {
  SlotName.value == 'IHaveName' ?
    SlotName.value = 'IHaveMony' :
    SlotName.value = 'IHaveName'
}

</script> 

<style lang='less' scoped>
.father_box {
  background-color: burlywood;
}
</style>