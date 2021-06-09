const environment = process.env.NODE_ENV || 'development'
console.log(process.env.BASE_URL);

const baseURL = process.env.BASE_URL || 'http://locahost:3000';
const apiHost = process.env.API_HOST || 'http://localhost:8000';
const apiPrefix = process.env.API_PREFIX || '/api';
const proxyPrefix = process.env.PROXY_PREFIX || '/api/';

export default {
    // RuntimeConfig
    publicRuntimeConfig: {
        baseURL,
    },
    // Global page headers: https://go.nuxtjs.dev/config-head
    head: {
        title: 'golang-front',
        htmlAttrs: {
            lang: 'jp'
        },
        meta: [
            { charset: 'utf-8' },
            { name: 'viewport', content: 'width=device-width, initial-scale=1' },
            { hid: 'description', name: 'description', content: '' }
        ],
        link: [
            { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
        ]
    },

    // Customize progress-bar
    loading: {
        color: 'white',
        height: '3px'
    },

    router: {
        extendRoutes(routes, resolve) {
            routes.push({
                name: 'error.404',
                path: '*',
                component: resolve('~/pages/error/404.vue'),
            })
        },
    },

    // Global CSS: https://go.nuxtjs.dev/config-css
    css: [
        '~/assets/css/app.css',
    ],

    // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
    plugins: [
        '~/plugins/axios',
        '~/plugins/api',
        '~/plugins/vee-validate',
    ],

    // Auto import components: https://go.nuxtjs.dev/config-components
    components: true,

    // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
    buildModules: [
        // https://go.nuxtjs.dev/typescript
        '@nuxt/typescript-build',
        // https://go.nuxtjs.dev/tailwindcss
        '@nuxtjs/tailwindcss',
    ],

    // Modules: https://go.nuxtjs.dev/config-modules
    modules: [
        '@nuxtjs/axios',
        '@nuxtjs/proxy',
    ],

    // Axios module configuration: https://go.nuxtjs.dev/config-axios
    axios: {
        prefix: apiPrefix,
        proxy: true,
        credentials: true,
    },

    proxy: {
        [proxyPrefix]: {
            target: apiHost,
        }
    },

    // Build Configuration: https://go.nuxtjs.dev/config-build
    build: {
        transpile: [
            "vee-validate/dist/rules", // For vee-validate
        ],
    }
}