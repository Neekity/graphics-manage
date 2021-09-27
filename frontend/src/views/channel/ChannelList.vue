<template>
  <v-app>

  </v-app>
</template>

<script>
export default {
  name: "ChannelList",
  data(){
    return{
      headers:[],
      channelItems:[],
    }
  },
  methods:{
    getChannel() {
      this.overlay += 1;
      this.$axios.post('/backend/material/channels')
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.channelItems = resData.data;
            } else {
              this.$toast('获取频道出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取频道出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
  }
}
</script>

<style scoped>

</style>