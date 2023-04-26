/// <reference types="vite/client" />
/// <reference types="vite-svg-loader" />


/**
 * 以下是定义接口出入参数
 */

declare namespace ApiRsp {
    type Common<T> = {
        code: number
        data: T,
        msg: string
    }
    type SendSMSInput = {
        phone_number: string
    }
    type LoginInput = {
        phone_number: string,
        login_type: string,
        code: number | string,
        password: string
    } 

    type UserResult = {
        id: number,
        avatar: string | null 
        name: string
        openid: null | string
        phone_number: string
        role_id: number
        unionid: null | string
    }
    type LoginResult= {
        user: UserResult,
        access_token: string,
        refresh_token: string,
    }
    type RoleResult = {
        name: string,
        code: string,
        description: string
    }

    type TheaterResult = {
        id: number,
        name: string,
        location: string,
        created_at: string,
        updated_at: string,
        count: number,
    }

    type HallResult = {
        id: number,
        name: string,
        theater_id: number,
        total_seats: number
    }
}


declare namespace App {
    type Pagination = {
        page_size: number,
        page_num: number
    }

    type IDBody = {
        id: number
    }
}