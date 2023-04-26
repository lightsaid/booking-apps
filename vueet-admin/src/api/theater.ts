import { http } from "../utils/http"

export function ListTheaters(arg: App.Pagination): Promise<ApiRsp.Common<ApiRsp.TheaterResult[]>> {
    return http.Get(`admin/theaters?page_size=${arg.page_size}&page_num=${arg.page_num}`, {})
}

export function GetTheater(id: number): Promise<ApiRsp.Common<ApiRsp.TheaterResult>> {
    return http.Get(`admin/theaters/${id}`, {})
}



export function ListHalls(id: number, arg: App.Pagination): Promise<ApiRsp.Common<ApiRsp.HallResult[]>> {
    return http.Get(`admin/halls?id=${id}&page_size=${arg.page_size}&page_num=${arg.page_num}`, {})
}
