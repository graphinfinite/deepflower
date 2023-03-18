
import axios from 'axios'

const API = axios.create()

API.defaults.baseURL = 'http://127.0.0.1:8787';

API.interceptors.request.use(function (config) {
    // Do something before request is sent

    console.log(`Bearer ${JSON.parse(localStorage.getItem("tokenAccess") || "")}`)

    config.headers.set('Content-Type', "application/json");
    config.headers.setAuthorization(`Bearer ${JSON.parse(localStorage.getItem("tokenAccess") || "")}`)
    return config;
  }, function (error) {
    // Do something with request error
    console.log(error)
    return Promise.reject(error);
  });

// Add a response interceptor
API.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    console.log(response.data) 
    return response;
  }, function (error) {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    return Promise.reject(error);
  });

export default API