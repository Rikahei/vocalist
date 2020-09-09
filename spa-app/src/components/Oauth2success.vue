<template>
  <v-container>
  <form>
    <input placeholder="paste your code here">
    <button>Submit2</button>
  </form>
  </v-container>
</template>

<script>
import Cookies from 'js-cookie';
import router from '../router.js'

export default {
  name: 'Oauth2success',
  data (){
    return {
      code: '',
    }
  },
  created() {
    console.log(this.$route.query.channelId);
    Cookies.set('channelId', this.$route.query.channelId);
    Cookies.set('title', this.$route.query.title);
    router.push({ path: '/' });
    location.reload();
  },
};
</script>