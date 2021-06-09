import Vue from 'vue'

import {
    extend,
    ValidationProvider,
    ValidationObserver,
    localize
} from 'vee-validate'
import ja from 'vee-validate/dist/locale/ja.json'
import { required, email, max } from 'vee-validate/dist/rules'

extend('required', required)
extend('email', email)
extend('max', max)

Vue.component('validation-provider', ValidationProvider)
Vue.component('validation-observer', ValidationObserver)

localize('ja', ja)
