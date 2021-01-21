<template>
  <v-container fluid>
    Search for miku w/ Docker with Region Code<br>
    <!-- <iframe class="centerPlayer" :src=url width="640px" height="360px" allow="autoplay;fullscreen"></iframe>  -->
    <!-- <iframe id="centerPlayer" :player="playStatus" width="560" height="315" :src=url frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe> -->
    <youtube :video-id="playerVideoId" :player-vars="playerVars" @ended="ended" class="centerPlayer"></youtube>
    <br>
    <v-row class="searchArea" >
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
                      width="340px"
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
        :list="myList"
        group="videos"
        style="height:200px"
      >
        <div
          class="list-group-item"
          v-for="(detail, index) in myList"
          :key="detail.videoId"
        >
          <!-- <v-col> -->
            <v-item>
              <v-card
                class="d-flex align-center"
                dark
                width="240px"
                height="100px"
                v-on:click="clickMyList(index, detail.videoId)"
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
import axios from 'axios';

export default {
  name: 'Home',
  props: ['details'],
  components: {
      draggable,
  },
  data (){
    return {
      playerVideoId: '7DVIoHWNOks',
      playerVars: {
        autoplay: 0
      },
      clickedid: '',
      list1: [],
      myList: [],
      playListSeq: 0,
    }
  },
  watch: { 
    'details': {
      handler(newVal) {
        // console.log('Prop changed: ', newVal, ' | was: ', oldVal)
        this.list1 = newVal;
      },
      deep: true
    },
    myList: function() {
      console.log('mylist updated')
      this.updateUserList();
    }
  },
  computed: {
    player() {
      return this.$refs.youtube.player
    }
  },

  methods: {
    clickVideo(videoId) {
      this.playerVideoId = videoId
      this.playerVars.autoplay = 1
    },
    clickMyList(seq, videoId) {
      this.playerVideoId = videoId
      this.playerVars.autoplay = 1
      this.playListSeq = seq
    },
    ended() {
      console.log('Video is ended')
      this.playNext()
    },
    playNext() {
      console.log('start next song')
      this.playListSeq += 1
      var nextSong = this.myList[this.playListSeq]
      if(typeof(nextSong) != "undefined"){
        this.playerVideoId = nextSong.videoId
      }else{
        console.log("This is last")
        this.playListSeq = 0
      }
    },
    updateUserList() {
      var myList = this.myList
      axios.post('/updateMyList', myList)
      .then((response) => {
        if(response.status == 200){
          console.log(response.data);
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
<style>

.centerPlayer{
  display: block;
  margin-left: auto;
  margin-right: auto;
}

label {
  padding:10px;
}
.searchArea{
  padding-bottom: 360px;
}
.myplaylist{
  padding: 6px;
  background-color:rgba(0, 0, 0, 0.8);
  position: fixed;
  Width: 100%;
  height: 20%;
  bottom: 0;
  overflow: scroll;
  z-index: 1;
}

</style>