
import axios from 'axios'

const API = axios.create()

API.defaults.baseURL = 'http://127.0.0.1:8787';

API.interceptors.request.use(function (config) {
    console.log(`Bearer ${JSON.parse(localStorage.getItem("tokenAccess") || "")}`)
    config.headers.set('Content-Type', "application/json");
    config.headers.setAuthorization(`Bearer ${JSON.parse(localStorage.getItem("tokenAccess") || "")}`)
    return config;
  }, function (error) {
    console.log(error)
    return Promise.reject(error);
  });

API.interceptors.response.use(function (response) {
    console.log(response.data) 
    return response;
  }, function (error) {
    return Promise.reject(error);
  });

export default API