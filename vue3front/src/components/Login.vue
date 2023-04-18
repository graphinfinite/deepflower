<template>

    <div id="login">
      <div id="description">
        <img src="@/assets/goat-icon-18-256.png" alt="goatlogo" width="32" height="32">
        <p> Glory to the goats!</p>
      </div>
      <div id="form">
        <form @submit.prevent="doLogin">
          <label for="username">username</label>
          <input type="text" id="username" v-model="username" placeholder="." autocomplete="off">
          <label for="password">password</label>&nbsp;
          <i class="fas" :class="[passwordFieldIcon]" @click="hidePassword = !hidePassword">*</i>
          <input :type="passwordFieldType" id="password" v-model="password" placeholder=".">

          <button type="submit">->...</button> 

          <div class="form-group">
          <div v-if="messageErr" class="alert alert-danger" role="alert">
            {{ messageErr }}
          </div>
        </div>
        </form>
        <br><br>
        <a id="signup" href="http://t.me/DeepFlowerbot">signup</a>
        
        
      </div>
    </div>
</template>

<script setup lang="ts">
  import { computed, ref } from "vue";
  import AuthService from "@/modules/auth"


  const messageErr = ref("")

  const username = ref("");
  const hidePassword = ref(true);
  const password = ref("");

  const isAuthorazed = ref(false)

  const passwordFieldIcon = computed(() => hidePassword.value ? "fa-eye" : "fa-eye-slash");
  const passwordFieldType = computed(() => hidePassword.value ? "password" : "text");
  const doLogin = () => AuthService.login(username.value,password.value)

</script>

<style scoped>

div#login {
  font-family: Verdana, sans-serif;
  align-items: center;
  background-color: #f0edf5;
  display: flex;
  justify-content: center;
  width: 100%;
  height: 100%;
}

div#login div#description {
  background-color: #bdb9c9;
  width: 200px;
  padding: 35px;
}

div#login div#description h1,
div#login div#description p {
  margin: 0;
}

div#login div#description p {
  font-size: 0.8em;
  color: #ffffff;
  margin-top: 10px;
}

div#login div#form {
  background-color: #000000;
  border-radius: 5px;
  box-shadow: 0px 0px 30px 0px #666;
  color: #7740f7;
  width: 260px;
  padding: 35px;
}



div#login div#form label,
div#login div#form input {
  outline: none;
  width: 100%;
}


div#login div#form i {

  cursor: pointer;
}


div#login div#form label {
  color: #6e5c7c;
  font-size: 1em;

}

div#login div#form input {
  background-color: transparent;
  color: #ffffff;
  font-size: 1em;
  margin-top: 10px;
  margin-bottom: 20px;
}

div#login div#form ::placeholder {
  color: #ecf0f1;
  opacity: 1;
}

div#login div#form button {
  background-color: #130227;
  cursor: pointer;
  border: none;
  padding: 10px;
  transition: background-color 1.5s ease-in-out;
  width: 100%;
}

div#login div#form button:hover {
  background-color: #000000;
}

@media screen and (max-width: 600px) {
  div#login {
    align-items: unset;
    background-color: unset;
    display: unset;
    justify-content: unset;
  }

  div#login div#description {
    margin: 0 auto;
    max-width: 350px;
    width: 100%;
  }

  div#login div#form {
    border-radius: unset;
    box-shadow: unset;
    width: 100%;
  }

  div#login div#form form {
    margin: 0 auto;
    max-width: 280px;
    width: 100%;
  }
}
</style>