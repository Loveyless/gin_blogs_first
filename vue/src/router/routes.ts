import type { RouteRecordRaw } from "vue-router";
import { createRouter, createWebHashHistory } from "vue-router";

// const metaRouters: any = import.meta.glob("./modules/*.ts", { eager: true });
// // * 处理路由
// export const AllRoute: RouteRecordRaw[] = [];
// Object.keys(metaRouters).forEach((item) => {
//   Object.keys(metaRouters[item]).forEach((key: any) => {
//     AllRoute.push(...metaRouters[item][key]);
//   });
// });

/**
 * @description 路由配置
 * @param path ==> 路由路径
 * @param name ==> 路由名称
 * @param redirect ==> 路由重定向
 * @param component ==> 路由组件
 * @param meta ==> 路由元信息
 * @param meta.noRequireAuth ==> 是否不要权限验证 默认需要
 * @param meta.noKeepAlive ==> 是否不缓存该路由 默认缓存
 * @param meta.title ==> 路由标题 目前用name
 * @param meta.key	==> 路由key 目前用不到
 * @param meta.noLoading ==> 页面loading 默认都有
 * @param meta.noNProgress ==> 顶部进度条 默认都有
 * */

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "登录",
    component: () => import("@v/login/index.vue"),
    meta: {
      noKeepAlive: true,
      noRequireAuth: true,
    },
  },
  {
    path: "/layout",
    component: () => import("@/layout/index.vue"),
    redirect: "/layout/home",
    children: [
      { path: "home", name: "主页", component: () => import("@v/Home/index.vue") },
      { path: "father_son", name: "父子之间", component: () => import("@v/FatherSon/Father.vue") },
      { path: "show_images", name: "图片插件+封装图片组件", component: () => import("@v/ShowImages/index.vue") },
      { path: "slot", name: "插槽", component: () => import("@v/Slot/FatherSlot.vue") },
      { path: "editor", name: "富文本", component: () => import("@v/WangEditor/index.vue") },
      { path: "list_wuhu", name: "动态菜单", component: () => import("@v/TestTemplate/ListWuhu!.vue") },
      { path: "test_demo", name: "pinia-setup和dialog组件", component: () => import("@v/TestDemo/index.vue") },
      { path: "provide_inject", name: "依赖注入", component: () => import("@v/Provide&Inject/index.vue") },
    ],
  },
  {
    path: "/:pathMatch(.*)*", //匹配没匹上的所有路由到404
    component: () => import("@v/404/index.vue"),
  },
];

//hash路由
const router = createRouter({
  history: createWebHashHistory(),
  routes,
  // 切换页面，滚动到最顶部
  scrollBehavior: () => ({ left: 0, top: 0 }),
});

export default router;
