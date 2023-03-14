import { defineStore } from 'pinia'
import storage from "@/utils/storage"

export const ProfileKey = "profile"

export const useProfileStore =  defineStore("profile", {
    state: ()=>({
        profile: {} as ApiRsp.LoginResult,
    }),
    getters: {

    },
    actions: {
        setProfile(payload: ApiRsp.LoginResult) {
            this.profile = payload
            storage.set(ProfileKey, payload)
        },
        logout() {
            this.profile = {} as ApiRsp.LoginResult
            storage.remove(ProfileKey)
        }
    }
})