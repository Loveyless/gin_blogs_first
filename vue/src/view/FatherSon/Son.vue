<template>
  <div class="son_box">
    son 组件 1

    <div>father传入的count{{ props.count }}</div>

    <el-button type="primary" @click="$emit('countAdd', 10)" plain>son button 直接使用emit</el-button>

    <el-button type="primary" @click="fatherCountAdd" plain>son button 声明emit再使用</el-button>

    <div>father 的 attrs {{ $attrs }}</div>
  </div>
  --------------------------------------------------
  <div v-bind="$attrs">
    son 组件 2

    <div>father 的 attrs {{ $attrs }}</div>

    多节点attr 这个节点继承father传入的 class 多节点继承样式会丢失 只能再son 组件中声明一个和穿透过来一样的 css样式
  </div>
</template>

<script lang="ts" setup>
//ts类型标注props
const props = defineProps<{
  count: number;
}>();

// 普通声明emit
// const emit = defineEmits(["countAdd"])

// ts声明
const emit = defineEmits<{
  (e: "countAdd", num: number): void;
  // (e: 'countAdd', num: number): void  不需要逗号 直接写
}>();

const fatherCountAdd = () => {
  emit("countAdd", 5);
};
</script>

<style lang="less" scoped>
.son_box {
  font-weight: 900;
}
</style>
