import axios from 'axios'



const API_URL = 'http://localhost:8080/auth/';

interface UserLogin {
    username: string,
    password: string
}

class AuthService {
  login(username:string, password:string) {
    return axios.post(API_URL + 'sign-in', {
        username: username,
        password: password
      })
      .then(response => {
        console.log(response.data)
        if (response.data.status === "ok") {
            localStorage.setItem('tokenAccess', JSON.stringify(response.data.token));
        }
        return response.data;
      });
  }
  logout() {
    localStorage.removeItem('tokenAccess');
  }
}

export default new AuthService();

