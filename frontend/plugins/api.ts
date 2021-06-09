import { NuxtAxiosInstance } from '@nuxtjs/axios'
import { Context, Inject } from '@nuxt/types/app'

export default function ({ $axios }: Context, inject: Inject) {
    const api = new ApiClient($axios)
    inject('api', api)
}

class ApiClient {

    constructor(
        private axios: NuxtAxiosInstance
    ) {
        //
    }

    public async getTest() {
        const res = await this.axios.get('message')
        console.log(res)
    }
}
