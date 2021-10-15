<template>
  <v-app>
    <message-show :content="message.content" :base-url="baseUrl" :subtitle="subtitle"
                  :title="title"></message-show>
    <v-overlay :value="overlay">
      <v-progress-circular
          indeterminate
          color="primary"
      ></v-progress-circular>
    </v-overlay>
  </v-app>
</template>

<script>

import MessageShow from "../../components/message/MessageShow";
export default {
  name: "MessageUserShow",
  components: {MessageShow},
  data(){
    return{
      overlay: false,
      title: '',
      subtitle: '',
      messageId:parseInt(this.$route.query.id),
      baseUrl:'/tinymce/',
      message: {},
    }
  },
  created() {
    this.getUserMessage();
  },
  methods: {
    getUserMessage() {
      if (this.messageId == 0) {
        return;
      }
      this.overlay = true;
      this.$graphicsHttp('post','/message/user/detail', {id: this.messageId})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.message = resData.data || {};
              this.title = this.message.title;
              this.subtitle = (this.message.author || '')+' '+this.message.send_time ;
            } else {
              this.$toast('获取消息出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取消息出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay = false);
    },
  },
}
</script>


<style>
.tox-tinymce {
  border:none !important
}
</style>
