import axios, { AxiosHeaders } from 'axios'

axios.defaults.baseURL = 'http://127.0.0.1:8787';

axios.interceptors.request.use((request) => {
  request.headers.set('Content-Type', "application/json");
  return request;
});

class AuthService {
  login(username:string, password:string) {
    let userLogin = JSON.stringify({
      username: username,
      password: password
    })
    console.log(userLogin)
    return axios.post('auth/sign-in', userLogin).then(response => {
        console.log(response.data)
        if (response.data.status === "ok") {

            console.log(response.data.data)
            localStorage.setItem('tokenAccess', JSON.stringify(response.data.data.token));
            location.reload()
        } 
        return response.data;
      });
  }
  logout() {
    localStorage.removeItem('tokenAccess');
    location.reload();
  }
}

export default new AuthService();

