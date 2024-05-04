import axios from 'axios'
import store from '@/store'
import { ElMessage } from 'element-plus'

// axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*'

const service = axios.create({
  baseURL:'/admin_api', // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 120000 // request timeout
})

service.interceptors.request.use(
  config => {
    // console.log('[request.js]: ', config)
    // do something before request is sent
    // if (store.getters.token) {
    //   // let each request carry token
    //   // ['X-Token'] is a custom headers key
    //   // please modify it according to the actual situation
    //   config.headers['Authorization'] = 'Bearer ' + getToken()
    // }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    // console.log(response)
    // const res = response.data

    // if the custom code is not 20000, it is judged as an error.
    return response
  },
  error => {
    let _msg = error.message
    const ignoreErrorList = ['DataSourceNotFound', 'DataSourceAlreadyExist']
    let _errorCode = ''
    if (error.response.data) {
      _msg = error.response.data.error ?? error.response.data
      _errorCode = error.response.data.code ?? ''
    }
    if (!_errorCode || ignoreErrorList.indexOf(_errorCode) === -1) {
      ElMessage({
        message: _msg,
        type: 'error',
        duration: 5 * 1000
      })
    }
    return Promise.reject(error.response.data)
  }
)

export default service
