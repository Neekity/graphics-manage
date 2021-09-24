<template>
  <v-app>
    <v-container fluid>
      <v-card width="300px" class="mx-auto my-12">
        <v-img v-show="!!message.cover_picture_url"
               :src="message.cover_picture_url"
               aspect-ratio="2.347"
        >
        </v-img>

        <v-card-title>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <span
                v-bind="attrs"
                v-on="on"
                class="d-inline-block text-truncate text-h6"
                style="max-width: 200px;"
              >
                 {{ message.title }}
              </span>
            </template>
            <span>{{ message.title }}</span>
          </v-tooltip>
        </v-card-title>
        <v-card-subtitle>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <span
                v-bind="attrs"
                v-on="on"
                class="d-inline-block text-truncate"
                style="max-width: 200px;"
              >
                 {{ message.abstract }}
              </span>
            </template>
            <span>{{ message.abstract }}</span>
          </v-tooltip>
        </v-card-subtitle>
      </v-card>
      <message-user-show v-if="showView" :content="message.content" :base-url="baseUrl" :subtitle="subtitle"
                         :title="title"></message-user-show>
      <v-card width="660px" class="mx-auto my-12">
        <v-card-text>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-chip
                v-bind="attrs"
                v-on="on"
                class="mr-2"
              >
                <v-icon left>
                  mdi-account-arrow-right-outline
                </v-icon>
                {{ message.sender_employee_name }}
              </v-chip>
            </template>
            <span>消息发送者</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-chip
                v-bind="attrs"
                v-on="on"
                class="mr-2"
              >
                <v-icon left>
                  mdi-radio-tower
                </v-icon>
                {{ message.channel }}
              </v-chip>
            </template>
            <span>消息所属频道</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-chip
                v-bind="attrs"
                v-on="on"
                class="mr-2"
              >
                <v-icon left>
                  mdi-clock-outline
                </v-icon>
                {{ message.send_at }}
              </v-chip>
            </template>
            <span>消息发送时间</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-chip
                v-bind="attrs"
                v-on="on"
                class="mr-2"
              >
                <v-icon left>
                  mdi-brightness-5
                </v-icon>
                {{ message.status }}
              </v-chip>
            </template>
            <span>消息状态</span>
          </v-tooltip>
          <v-divider></v-divider>
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
                mdi-account-arrow-left-outline
              </v-icon>
              {{ node }}
            </v-chip>
          </v-scroll-x-transition>
        </v-card-text>
      </v-card>
    </v-container>
    <v-overlay :value="overlay">
      <v-progress-circular
        indeterminate
        color="primary"
      ></v-progress-circular>
    </v-overlay>
  </v-app>
</template>

<script>
import MessageUserShow from "./MessageUserShow";

export default {
  name: "MessageOwnerShow",
  components: {MessageUserShow},
  props: ['messageId', 'baseUrl'],
  data() {
    return {
      overlay: false,
      showView: false,
      title: '',
      subtitle: '',
      message: {},
    }
  },
  created() {
    this.getMaterial();
  },
  methods: {
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
            this.showView = true;
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
