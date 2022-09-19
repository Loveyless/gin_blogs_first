import { defineStore, createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate";

// 使用setup模式定义
export const GlobalStore = defineStore(
  "GlobalStore",
  () => {
    let Id = ref<string>("");
    let Username = ref<string>("");

    function setUsername(value: string) {
      Username.value = value;
    }

    return { Id, Username, setUsername };
  },
  {
    persist: true,
  }
);

// 使用options API模式定义
// export const GlobalStore = defineStore("GlobalStore", {
//   // 定义state
//   state: () => ({
//     Id: "",
//     Username: "",
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
//       return state.Id;
//     },
//   },
//   persist: true,
// });

const pinia = createPinia();
//注册持久化插件
pinia.use(piniaPluginPersistedstate);
export default pinia;
