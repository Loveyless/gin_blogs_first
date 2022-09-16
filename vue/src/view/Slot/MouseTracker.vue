<template>
  <slot :x="x" :y="y" />
</template>

<script lang='ts' setup>

//hooks XY
const { x, y } = useMouse()





// 这里的获取坐标抽离出去换了 hooks

// const x = ref<number>(0)
// const y = ref<number>(0)

// const update = (e: any) =>  {
//   x.value = e.pageX
//   y.value = e.pageY
// }

// onMounted(() => window.addEventListener('mousemove', update))
// onUnmounted(() => window.removeEventListener('mousemove', update))

</script> 
