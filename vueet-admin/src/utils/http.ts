import axios, { AxiosRequestConfig } from "axios"
import { ElMessage } from "element-plus"
import storage from "@/utils/storage"
import { ProfileKey } from "@/store/profile"

const defaultConfig = {
    timeout: 10000,
    baseURL: import.meta.env.VITE_BASE_URL
}

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
                if (response.data.code === 1000) {
                    return response
                }else{
                    ElMessage.error(response.data.msg)
                    return Promise.reject(response)
                }
            },
            err => {
                ElMessage.error(err?.response?.data?.msg || err.message)
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
}

export const http = new Http()