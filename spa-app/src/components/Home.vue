<template>
  <v-container fluid>
    Search for miku w/ Docker<br>
    <iframe width="560" height="315" :src=url frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
    <br>
    <v-row>
      <v-col cols="12">
        <v-item-group>
          <v-row>
            <draggable 
              class="row dragArea list-group"
              :list="list1"
              :group="{ name: 'videos', pull: 'clone', put: false }"
            >
              <div
                v-for="detail in details"
                :key="detail.videoId"
              >
                <v-col>
                  <v-item>
                    <v-card
                      class="d-flex align-center"
                      dark
                      height="200px"
                      max-width="340px"
                      v-on:click="clickVideo(detail.videoId)"
                    >
                      <v-img
                        class="white--text align-end"
                        height="200px"
                        :src = detail.thumbnail
                      >
                        <v-card-title class="subtitle-1">{{detail.title}}</v-card-title>
                      </v-img>
                    </v-card>
                  </v-item>
                </v-col>
              </div>

            </draggable>
          </v-row>
        </v-item-group>
      </v-col>
    </v-row>
    
    <v-row class="myplaylist">
      <label>This is playlist</label>
      <draggable
        class="dragArea list-group row"
        :list="list2"
        group="videos"
        style="height:200px"
      >
        <div
          class="list-group-item"
          v-for="detail in list2"
          :key="detail.videoId"
        >
          <!-- <v-col> -->
            <v-item>
              <v-card
                class="d-flex align-center"
                dark
                width="240px"
                height="100px"
                v-on:click="clickVideo(detail.videoId)"
              >
                <v-img
                  class="white--text align-end"
                  height="100px"
                  :src = detail.thumbnail
                >
                  <v-card-title class="subtitle-1">{{detail.title}}</v-card-title>
                </v-img>
              </v-card>
            </v-item>
          <!-- </v-col> -->
        </div>
      </draggable>
    </v-row>
  </v-container>
</template>

<script>
import draggable from 'vuedraggable';
export default {
  name: 'Home',
  props: ['details'],
  components: {
      draggable,
  },
  data (){
    return {
      url: 'https://www.youtube.com/embed/AufydOsiD6M',
      clickedid: '',
      list1: [],
      list2: [],
    }
  },
  watch: { 
    'details': {
      handler(newVal, oldVal) {
        console.log('Prop changed: ', newVal, ' | was: ', oldVal)
        this.list1 = newVal;
      },
      deep: true
    },
  },
  methods: {
    clickVideo(videoId) {
      var sUrl = videoId
      console.log (sUrl);
      this.url = "https://www.youtube.com/embed/" + sUrl
    }
  }
};
</script>
<style>
label {
  padding:10px;
}
.myplaylist{
  background-color:rgba(140, 140, 140, 0.6);
  position: fixed;
  Width: 100%;
  height: 20%;
  bottom: 0;
  overflow: scroll;
}

</style>
