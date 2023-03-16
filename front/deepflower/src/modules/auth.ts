import { reactive, toRefs } from 'vue'



// https://dev.to/adamcowley/how-to-build-an-authentication-into-a-vue3-application-200b
interface User {
    id: string;
    email: string;
    dateOfBirth: Date;
    firstName: string;
    lastName: string;
    access_token: string;
}

interface UserState {
    authenticating: boolean;
    user?: User;
    error?: Error;
}

interface AuthState {
    authenticating: boolean;
    user?: User;
    error?: Error;
}


const state = reactive<AuthState>({
    authenticating: false,
    user: undefined,
    error: undefined,
})

const AUTH_KEY = ""

export const useAuth = () => {
    const setUser = (payload: User, remember: boolean) => {
      if ( remember ) {
        // Save
        window.localStorage.setItem(AUTH_KEY, payload[ AUTH_TOKEN ])
      }
  
      state.user = payload
      state.error = undefined
    }
  
    const logout = (): Promise<void> => {
      window.localStorage.removeItem(AUTH_KEY)
      return Promise.resolve(state.user = undefined)
    }
  
    return {
      setUser,
      logout,
      ...toRefs(state), // authenticating, user, error
    }
  }


