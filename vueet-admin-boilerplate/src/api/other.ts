import { http } from "../utils/http"

export function SendSMS(data: ApiRsp.SendSMSInput): Promise<ApiRsp.Common<number>>{
    return http.Post("/sms", {data})
}