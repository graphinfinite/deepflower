
import axios from 'axios'

const api = axios.create({
    baseURL: `127.0.0.1:8080/`

})
api.interceptors.request.use(function (config) {
    // Do something before request is sent
    config.headers.setAuthorization(`Bearer ${localStorage.getItem("accessToken")}`)
    return config;
  }, function (error) {
    // Do something with request error
    console.log(error)
    return Promise.reject(error);
  });

// Add a response interceptor
axios.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    console.log(response.data.status) 
    return response;
  }, function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
  });

export default api