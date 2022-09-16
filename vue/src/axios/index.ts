import type { AxiosInstance, AxiosError, AxiosRequestConfig, AxiosResponse } from "axios";
import instance from "@/axios/config";
import { openLoading, closeLoading } from "@/hooks/loading";
import NProgress from "@/hooks/nprogress";
import { Loading } from "@element-plus/icons-vue";
import { AxiosCanceler } from "@/axios/cancel";

//实例化取消请求
const axiosCanceler = new AxiosCanceler();

//拦截
instance.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    const globalStore = GlobalStore();
    // 将当前请求添加到 pending 中
    axiosCanceler.addPending(config);

    //全局loading 和 进度条 请求头中 { loaind : true , loadingdark : true , nprogress : true }
    config.headers!.nprogress && NProgress.start(); //默认需要进度条
    if (!config.headers!.loading) {
      //默认不要loading
      if (config.headers!.loadingdark) {
        openLoading();
      } else {
        openLoading(true); //白色loading
      }
    }

    return { ...config };
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();
  }
);

//响应
instance.interceptors.response.use(
  (response: AxiosResponse) => {
    const { data, config } = response;
    // 在请求结束后，移除本次请求
    axiosCanceler.removePending(config);

    closeLoading();
    NProgress.done();
    return response;
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();
  }
);

export const http = instance;
