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

    //全局loading 和 进度条 请求头中 { loading : true , loadingdark : true , nprogress : true }
    config.headers!.nprogress && NProgress.start(); //默认不需要进度条
    if (config.headers!.loading) {
      //默认不要loading
      if (config.headers!.loadingdark) {
        openLoading();
      } else {
        openLoading(true); //白色loading
      }
    }

    // 打印请求
    console.log("-------------------");
    console.log("地址", config.method, config.url);
    console.log("请求头", config.headers);
    console.log("请求参数", config.data);

    return { ...config };
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();

    //前置请求错误
    console.log("前置请求错误", err.message);
    return err;
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

    console.log("请求结果", response.data, typeof response.data);
    return response;
    1;
  },
  (err: AxiosError) => {
    closeLoading();
    NProgress.done();

    //后置请求错误
    console.log("后置请求错误", err);
    console.log("后置错误详情", err.response);
    return err;
  }
);

export const http = instance;
