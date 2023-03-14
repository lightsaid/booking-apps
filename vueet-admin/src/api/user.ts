import { http } from "../utils/http"

export function Login(data: ApiRsp.LoginInput): Promise<ApiRsp.Common<ApiRsp.LoginResult>> {
   return http.Post("/auth/login", {data})
}

export function GetProfile(): Promise<ApiRsp.Common<ApiRsp.UserResult>> {
    return http.Post("/admin/profile", {})
}