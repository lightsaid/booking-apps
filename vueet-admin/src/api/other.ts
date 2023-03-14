import { http } from "../utils/http"

export function SendSMS(data: ApiRsp.SendSMSInput): Promise<ApiRsp.Common<number>>{
    return http.Post("/sms", {data})
}

export function GetRoleById(id: number): Promise<ApiRsp.Common<ApiRsp.RoleResult>>{
    return http.Get(`/admin/roles/${id}`, {})
}