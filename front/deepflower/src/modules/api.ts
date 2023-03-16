
import axios from 'axios'

const api = axios.create({
    baseURL: `127.0.0.1:8080/`,
    headers: {
        Authorization: `Bearer ${localStorage.getItem("accessToken")}`
      }
  })

export default api

  //api.interceptors.request