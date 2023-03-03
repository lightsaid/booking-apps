import axios, { AxiosRequestConfig } from "axios"

const defaultConfig = {
    timeout: 10000,
    baseUrl: ''
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
                return response
            },
            err => {
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
        return Http.axiosInstance.get(url, params).then(res=>res.data).catch()
    }
}

export const http = new Http()