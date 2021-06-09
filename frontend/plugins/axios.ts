import { Context } from '@nuxt/types/app'
import { AxiosError, AxiosRequestConfig } from 'axios';


export default function ({ $axios, error }: Context) {
    $axios.onRequest((config: AxiosRequestConfig) => {
        console.log('request to ' + config.url)
    })

    $axios.onError((axiosError: AxiosError) => {
        const response = axiosError.response
        const status = response?.status ?? 500

        error({ statusCode: status })
    })
}
