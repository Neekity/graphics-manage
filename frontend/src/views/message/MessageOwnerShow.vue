<template>
  <v-app>
    <v-card width="660px" class="mx-auto my-12">
      <v-card-title>
        <v-card>
          <v-img
            :src="message.cover_picture_url"
            aspect-ratio="2.347"
          >
          </v-img>
          <v-card-title class="text-h6">
            {{ message.title }}
          </v-card-title>
          <v-card-subtitle>
            {{ message.abstract }}
          </v-card-subtitle>
          <v-card-subtitle>
            {{ message.send_at }}
          </v-card-subtitle>
        </v-card>
      </v-card-title>
      <v-card-text>
        <v-row
          align="center"
          class="mx-0 tinymce-content"
        >
          <message-user-show :content="message.content" :skin-url="skinUrl" :subtitle="subtitle" :title="title"></message-user-show>
        </v-row>
      </v-card-text>
      <v-divider class="mx-4"></v-divider>
      <v-card-subtitle>
          <v-list>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>{{ message.sender }}</v-list-item-title>
                <v-list-item-subtitle>消息发送者</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title>{{ message.channel}}</v-list-item-title>
                <v-list-item-subtitle>消息所属频道</v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-list-item>
              <v-list-item-content>
                <v-scroll-x-transition
                  group
                  hide-on-leave
                >
                  <v-chip
                    v-for="(node, i) in message.receivers"
                    :key="i"
                    color="primary"
                    outlined
                    small
                    class="ma-1"
                  >
                    <v-icon
                      left
                      small
                    >
                      mdi-account
                    </v-icon>
                    {{ node.name }}
                  </v-chip>
                </v-scroll-x-transition>
              </v-list-item-content>
            </v-list-item>
          </v-list>
      </v-card-subtitle>
      <v-divider class="mx-4"></v-divider>
    </v-card>
    <v-overlay :value="overlay">
      <v-progress-circular
        indeterminate
        color="primary"
      ></v-progress-circular>
    </v-overlay>
  </v-app>
</template>

<script>
import MessageShow from "./MessageShow";

export default {
  name: "MessageOnwerShow",
  components: {MessageShow},
  props:['messageId','skinUrl'],
  data(){
    return {
      overlay: false,
      title: '',
      subtitle: '',
      message:{},
    }
  },
  created() {
    this.getMaterial();
  },
  methods:{
    getMaterial() {
      if (this.messageId == 0) {
        return;
      }
      this.overlay = true;
      this.$http.post('/graphics/message/owner/data', {id: this.messageId})
        .then(response => {
          let resData = response.data;
          if (resData.code === 0) {
            this.message = resData.data.message;
            this.title = resData.data.title;
            this.subtitle = resData.data.subtitle;

          } else {
            this.$toast('获取消息出错：' + resData.message, {
              type: 'error',
            });
          }
        })
        .catch(error => {
          console.log(error);
          this.$toast('保存出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
  },
}
</script>
