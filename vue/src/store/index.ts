import { defineStore, createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

// 使用setup模式定义
export const GlobalStore = defineStore(
  "GlobalStore",
  () => {
    let token = ref<string>("");

    let count = ref<number>(0);
    function countAdd(): void {
      count.value++;
    }

    function doubleCount(): number {
      return count.value * 2;
    }

    return { token, count, countAdd, doubleCount };
  },
  {
    persist: true,
  }
);

const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
export default pinia;

// 使用options API模式定义
// export const Global = defineStore("Global", {
//   // 定义state
//   state: () => ({
//     count1: 1,
//   }),
//   // 定义action
//   actions: {
//     increment() {
//       this.count1++;
//     },
//   },
//   // 定义getters
//   getters: {
//     doubleCount(state) {
//       return state.count1 * 2;
//     },
//   },
// });
