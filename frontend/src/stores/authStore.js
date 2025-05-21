import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth',  {
    state: () => ({
        user: null
    }),

    actions: {
        async getUser() {
            try {
                const res = await axios.get(
                    '/api/v1/auth/me',
                )
                this.user = res.data.payload
            } catch {
                this.user = null
            }
        },
    }
})