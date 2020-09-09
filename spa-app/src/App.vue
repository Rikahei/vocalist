<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    <v-btn><router-link to="/">Return</router-link></v-btn>
    <v-btn v-if="username" v-on:click="logout">Logout</v-btn>
    <v-btn v-else><router-link to="/login">Login</router-link></v-btn>
    <v-text-field v-model="keyword" solo hide-details single-line>
    </v-text-field>
    <v-btn v-on:click="search">search</v-btn>
    <v-spacer></v-spacer>
    <span icon v-if="username">{{ username }}</span>
    </v-app-bar>

    <v-content>
      <router-view :details="details"></router-view>
    </v-content>
  </v-app>
</template>

<script>
import Cookies from 'js-cookie';
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      username: Cookies.get('title'),
      ytUrl: 'https://www.youtube.com/embed/',
      keyword: '',
      url: 'https://www.youtube.com/embed/AufydOsiD6M',
      details: '',
    }
  },
  methods: {
    logout: function () {
      Cookies.remove('title');
      Cookies.remove('channelId');
      location.reload();
    },
    search(e) {
      e.preventDefault();
      const params = new URLSearchParams();
      params.append('keyword', this.keyword);
      axios.post('/search', params)
      .then((response) => {
        if(response.status == 200){
          var sUrl = Object.keys(response.data)[0];
          this.url = "https://www.youtube.com/embed/" + sUrl;
          var arr = [];
          var dj = Object.values(response.data);
          dj.forEach(function(detail, key){
            arr[key] = JSON.parse(detail);
          })
          console.log(arr);
          this.details = arr;
        }
      })
      .catch(function (error) {
        alert("error!")
        console.log(error);
      });
    },
  }
};
</script>
