<template>
  <v-container>
  <form @submit="formSubmit">
    <input type="image" src="./btn_google_signin_light_normal_web_2x.png" border="0" alt="Submit" />
  </form>
  <br>
  <h3>{{ welcomeMsg }}</h3>
  </v-container>
</template>

<script>
import axios from 'axios';
import Cookies from 'js-cookie';
// import router from '../router.js'

export default {
  name: 'Login',

  data (){
    return {
      oauth2token: '',
    }
  },
  computed: {
    welcomeMsg: function() {
      if(document.cookie){
        return 'Welcome come back ' + Cookies.get('title');
      }else{
        return 'Please login';
      }
    }
  },
  methods: {
    formSubmit(e) {
      e.preventDefault();
      const params = new URLSearchParams();
      params.append('username', this.username);
      params.append('password', this.password);
      axios.get('/oauth2', params)
      .then(function (response) {
        if(response.status == 200){
          // Cookies.set('Username', response.data.username);
          // Cookies.set('Token', response.data.token, { expires: 1 });
          // router.push({ path: '/' });
          // location.reload();
          console.log(response.data);
          window.location.href = response.data;
        }
      })
      .catch(function (error) {
        alert("Login error!")
        console.log(error);
      });
    },
    codeSubmit(e) {
      e.preventDefault();
      const params = new URLSearchParams();
      params.append('token', this.oauth2token);
      axios.post('/oauth2token', params)
      .then(function (response) {
        if(response.status == 200){
          console.log(response.data);
          
        }
      })
    }
  }
};
</script>

<style scoped>
input {
  border-style: solid;
  border-color: grey;
}
</style>
