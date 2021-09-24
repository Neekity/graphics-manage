<template>
  <v-app>
    <v-card width="660px" class="mx-auto my-12">
      <v-card-title>
        <div class="text-h5">定制频道听众</div>
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12"
                 lg="6"
                 md="6"
                 sm="6"
                 xs="12">
            <v-select
              v-model="channel"
              item-text="name"
              item-value="id"
              :items="channelItems"
              label="频道"
              :rules="[value => !!value || '请选择定制听众的频道！',]"
              :disabled="disableChannel"
              dense
              outlined
            ></v-select>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="search"
              label="输入关键字搜索听众"
              clearable
              dense
              outlined
              @input="handleSearch"
              clear-icon="mdi-close-circle-outline"
            ></v-text-field>
            <v-treeview
              ref="items"
              v-model="selection"
              selection-type="independent"
              :items="items"
              selectable
              return-object
              dense
              hoverable
              :open.sync="open"
              item-text="text"
              :search="search"
            ></v-treeview>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col>
            <v-card-text>
              <v-scroll-x-transition
                group
                hide-on-leave
              >
                <v-chip
                  v-for="(node, i) in selection"
                  :key="i"
                  color="primary"
                  outlined
                  small
                  close
                  class="ma-1"
                  @click:close="selection.splice(i,1)"
                >
                  <v-icon
                    left
                    small
                  >
                    {{ node.type == 'employee' ? ' mdi-account' : 'mdi-account-multiple' }}
                  </v-icon>
                  {{ node.text }}
                </v-chip>
              </v-scroll-x-transition>
            </v-card-text>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-btn
          color="success"
          @click="storeChannelAudience"
          :disabled="!channel"
          class="text-center"
        >
          保存
        </v-btn>
      </v-card-actions>
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

export default {
  name: "SelectChannelAudience",
  props: ['channelId'],
  data() {
    return {
      open: [],
      allOpened: false,
      disableChannel: false,
      lastOpen: [],
      overlay: false,
      channel: '0',
      channelItems: [],
      search: null,
      selection: [],
      items: [],
    }
  },
  watch: {
    channel: {
      handler() {
        if (this.channel != 0) {
          this.getChannelAudience()
        }
      },
    },
  },
  created() {
    this.getChannels();
    this.getAllAudience();
    this.channel = this.channelId;
    if (this.channel != 0) {
      this.disableChannel = true;
    }
  },
  methods: {
    getChannels() {
      this.overlay = true;
      this.$http.post('/graphics/manage/channel/channels')
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
          this.$toast('保存出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
    getChannelAudience() {
      this.overlay = true;
      this.$http.post('/graphics/manage/channel/audience/data', {channel_id: this.channel})
        .then(response => {
          let resData = response.data;
          if (resData.code === 0) {
            this.selection = resData.data;
          } else {
            this.$toast('获取频道听众出错：' + resData.message, {
              type: 'error',
            });
          }
        })
        .catch(error => {
          console.log(error);
          this.$toast('获取频道听众出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
    getAllAudience() {
      this.overlay = true;
      this.$http.post('/graphics/manage/channel/employees')
        .then(response => {
          let resData = response.data;
          if (resData.code === 0) {
            this.items = resData.data;
          } else {
            this.$toast('获取所有听众出错：' + resData.message, {
              type: 'error',
            });
          }
        })
        .catch(error => {
          console.log(error);
          this.$toast('出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
    storeChannelAudience() {
      this.overlay = true;
      let receivers = [];
      this.selection.forEach(function (element) {
        receivers.push({id: element.id, type: element.type, name: element.text})
      });
      this.$http.post('/graphics/manage/channel/audience/store', {channel_id: this.channel, audiences: receivers})
        .then(response => {
          let resData = response.data;
          if (resData.code === 0) {
            this.$toast("保存成功！", {
              type: 'success',
              timeout: 2000
            });
            location.href = '/graphics/manage/channel';
          } else {
            this.$toast('保存频道听众出错：' + resData.message, {
              type: 'error',
            });
          }
        })
        .catch(error => {
          console.log(error);
          this.$toast('保存频道听众出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
    handleSearch: function (val) {
      if (val) {
        if (!this.allOpened) {
          this.lastOpen = this.open;
          this.allOpened = true;
          this.$refs.items.updateAll(true);
        }
      } else {
        this.$refs.items.updateAll(false);
        this.allOpened = false;
        this.open = this.lastOpen;
      }
    }
  },
}
</script>
