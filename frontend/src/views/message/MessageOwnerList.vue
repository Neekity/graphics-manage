<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row class="pl-2">
          <v-col
              cols="12"
              lg="3"
              md="4"
              sm="5"
              xs="12"
          >
            <v-text-field
                v-model="title"
                label="请输入标题"
                single-line
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getOwnerMessage"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-col
              cols="12"
              lg="3"
              md="4"
              sm="5"
              xs="12"
          >
            <v-select
                v-model="channel"
                item-text="name"
                item-value="id"
                :items="channelItems"
                label="选择频道"
                dense
                outlined
            ></v-select>
          </v-col>

          <v-spacer></v-spacer>
          <v-tooltip bottom v-if="showTable">
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                  v-bind="attrs"
                  v-on="on"
                  icon
                  color="blue darken-2"
                  @click="showTable=false"
              >
                <v-icon>mdi-image-multiple</v-icon>
              </v-btn>
            </template>
            <span>切换成图文卡片显示</span>
          </v-tooltip>
          <v-tooltip bottom v-if="!showTable">
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                  v-bind="attrs"
                  v-on="on"
                  icon
                  color="blue darken-2"
                  @click="showTable=true"
              >
                <v-icon>mdi-table</v-icon>
              </v-btn>
            </template>
            <span>切换成表格显示</span>
          </v-tooltip>
        </v-row>

      </v-card-title>
      <v-card-text>
        <v-data-table
            v-if="showTable"
            :headers="headers"
            :items="messages"
            item-key="id"
            class="elevation-0"
            :footer-props="{
              showFirstLastPage: true,
              firstIcon: 'mdi-arrow-collapse-left',
              lastIcon: 'mdi-arrow-collapse-right',
              prevIcon: 'mdi-minus',
              nextIcon: 'mdi-plus'
            }"
        >
          <template v-slot:item.actions="{ item }">
            <v-btn
                class="mr-2"
                color="blue darken-2"
                icon
                @click="view(item.id)"
            ><v-icon>
              mdi-eye
            </v-icon></v-btn>
            <v-btn
                class="mr-2"
                color="red darken-2"
                icon
                @click="confirm(item.id)"
            ><v-icon>
              mdi-delete
            </v-icon></v-btn>
          </template>
          <template v-slot:no-data>
            <v-btn
                color="primary"
                @click="initialize"
            >
              Reset
            </v-btn>
          </template>
          <template v-slot:item.title="{ item }">
            <TextTruncate :text="item.title"></TextTruncate>
          </template>
          <template v-slot:item.abstract="{ item }">
            <TextTruncate :text="item.abstract"></TextTruncate>
          </template>
        </v-data-table>
        <v-lazy
            v-if="!showTable"
            v-model="isActive"
            :options="{
          threshold: .5
        }"
            min-height="200"
            transition="fade-transition"
        >
          <v-container fluid>
            <v-row>
              <v-col
                  v-for="message in messages"
                  :key="message.id"
                  cols="12"
                  lg="3"
                  md="4"
                  sm="6"
                  xs="12"
              >
                <v-card @click="view(message.id)">
                  <v-img
                      :src="message.cover_picture_url"
                      aspect-ratio="2.347"
                  >
                  </v-img>
                  <v-card-title class="text-h6">
                    {{ message.title }}
                  </v-card-title>
                  <v-card-subtitle class="text-h6">
                    {{ message.abstract }}
                  </v-card-subtitle>
                  <v-divider></v-divider>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                        class="mr-2"
                        color="blue darken-2"
                        icon
                        @click="view(item.id)"
                    ><v-icon>
                      mdi-eye
                    </v-icon></v-btn>
                    <v-btn
                        class="mr-2"
                        color="red darken-2"
                        icon
                        @click="confirm(item.id)"
                    ><v-icon>
                      mdi-delete
                    </v-icon></v-btn>
                  </v-card-actions>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-lazy>
      </v-card-text>
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
import TextTruncate from "../../components/TextTruncate";
export default {
  name: 'MessageMasonry',
  components: {TextTruncate},
  data() {
    return {
      isActive: false,
      messages: [],
      overlay: false,
      title: "",
      channel: 0,
      showTable:false,
      channelItems: [],
      headers: [{
        text: '标题',
        align: 'start',
        value: 'title',
      },
        {text: '摘要', value: 'abstract', sortable: false},
        {text: '作者', value: 'author', sortable: false},
        {text: '发送时间', value: 'send_time', sortable: false},
        {text: '操作', value: 'actions', sortable: false},
      ],
    }
  },
  created() {
    this.getOwnerMessage();
    this.getChannels()
  },
  watch: {
    channel: function () {
        this.getOwnerMessage()
    }
  },
  methods: {
    getOwnerMessage() {
      this.overlay += 1;
      this.$graphicsHttp('post', 'message/owner/list', {
        title: this.title, channel_id: this.channel
      }).then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.messages = resData.data || [];
            } else {
              this.$toast('获取消息列表出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取消息列表出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    getChannels() {
      this.overlay += 1;
      this.$graphicsHttp('post', 'material/channels')
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.channelItems = resData.data || [];
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
    view(id) {
      location.href = '/message-owner/show?id=' + id;
    },
  },
}
</script>
