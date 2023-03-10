<template>

  <div id="left" > </div>
    <div id="right">

      <form class="login-form" @submit.prevent="login">
        <h2>Sign in</h2>
        <hr>
        <label>Username </label>
        <input required v-model="username" type="text" placeholder="..." />
        <hr>
        <label> Password </label>
        <input
          required
          v-model="password"
          type="password"
          placeholder="..."
        />

        <button type="submit">Login</button>
      </form>
      <a href="http://t.me/DeepFlowerbot">SignUpTelegram</a>
    </div>
  </template>
  
<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;

#left { position: absolute; left: 0; top: 0; width: 20%; }
#right { position: absolute; right: 0; top: 0; width: 70%; }
  </style>
  
  <script  lang="ts">
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
       // this.$store.dispatch("/auth/login", { username, password }).then(() => {
        //  this.$router.push("/");
       // });
      }
    }
  };
  </script>