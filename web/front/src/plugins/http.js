import Vue from 'vue'
import axios from 'axios'

// axios.defaults.baseURL = 'http://localhost:3000/api/v1'
//服务器环境
axios.defaults.baseURL = 'http://20.51.164.91:3000/api/v1'
Vue.prototype.$http = axios
