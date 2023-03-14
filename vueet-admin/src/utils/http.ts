import axios, { AxiosRequestConfig } from "axios"
import { ElMessage } from "element-plus"
import storage from "@/utils/storage"
import { ProfileKey, useProfileStore } from "@/store/profile"

const defaultConfig = {
    timeout: 10000,
    baseURL: import.meta.env.VITE_BASE_URL
}

const successCodes = [200, 201, 202, 203, 204, 205]

class Http {
    constructor() {
        this.interceptorsRequest()
        this.interceptorsResponse()
    }

    private static axiosInstance = axios.create(defaultConfig)

    // 请求拦截
    private interceptorsRequest() {
        Http.axiosInstance.interceptors.request.use(
            config => {
                const data = storage.get(ProfileKey)
                config.headers.Authorization = `Bearer ${data?.access_token}`
                return config
            },
            err => {
                return Promise.reject(err)
            }
        )
    }

    // 响应拦截
    private interceptorsResponse() {
        Http.axiosInstance.interceptors.response.use(
            response => {
                if (successCodes.includes(response.data.code)) {
                    return response
                }else{
                    ElMessage.error(response.data.msg)
                    return Promise.reject(response)
                }
            },
            err => {
                ElMessage.error(err?.response?.data?.msg || err.message)
                if (err?.response?.data?.code == 401) {
                     const store =  useProfileStore()
                    store.logout()
                }
                return Promise.reject(err)
            }
        )
    }

    // Get 请求
    public Get<T>(url: string, params: AxiosRequestConfig):Promise<T> {
        return Http.axiosInstance.get(url, params).then(res=>res.data).catch()
    }

    // Post 请求
    public Post<T>(url: string, params: AxiosRequestConfig):Promise<T> {
        return Http.axiosInstance.post(url, params.data).then(res=>res.data).catch()
    }

     // Post 请求
     public Put<T>(url: string, params: AxiosRequestConfig):Promise<T> {
        return Http.axiosInstance.put(url, params.data).then(res=>res.data).catch()
    }
}

export const http = new Http()