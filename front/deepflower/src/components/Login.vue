<template>
    <div>
      <form class="login-form" @submit.prevent="login">
        <h2>Sign in</h2>
        <label>Username </label>
        <input required v-model="username" type="text" placeholder="..." />
        <label> Password </label>
        <input
          required
          v-model="password"
          type="password"
          placeholder="..."
        />
        <hr />
        <button type="submit">Login</button>
      </form>
      <hr />
      <a href="http://t.me/DeepFlowerbot">SignUpTelegram</a>
    </div>
  </template>
  
  <style>
  .login {
    display: flex;
    flex-direction: row;
    width: 200px;
    padding: 10px;
  }
  </style>
  
  <script>
  export default {
    name: "login-form",
    data() {
      return {
        username: "",
        password: ""
      };
    },
    methods: {
      login: function() {
        const { username, password } = this;
        this.$store.dispatch("/auth/login", { username, password }).then(() => {
          this.$router.push("/");
        });
      }
    }
  };
  </script>