import { reactive, ref } from "vue"
import { ListTheaters, ListHalls, GetTheater } from "@/api/theater"

export const useTheater = () => {
    const theaters = reactive({
        data: [] as ApiRsp.TheaterResult[],
        page_num: 1,
        page_size: 10,
        total: 0
    })

    const theater = ref<ApiRsp.TheaterResult>()

    const getTheaters = (page_num: number, page_size: number) => {
        theaters.page_num = page_num
        theaters.page_size = page_size
        ListTheaters({page_num: page_num, page_size: page_size}).then(res=> {
            theaters.data = res.data
            theaters.total = res.data[0]?.count
        })
    }

    const getTheater = (id: number) => {
        GetTheater(id).then(res=> {
            theater.value = res.data
        })
    } 

    return {
        theater,
        getTheater,
        theaters,
        getTheaters
    }
}


export const useHall = () => {
    const halls = reactive({
        data: [] as ApiRsp.HallResult[]
    })

    const getHalls = (id: number,page_num: number, page_size: number) => {
        ListHalls(id, {page_num: page_num, page_size: page_size}).then(res=> {
            halls.data = res.data
        })
    }

    return {
        halls,
        getHalls
    }
}