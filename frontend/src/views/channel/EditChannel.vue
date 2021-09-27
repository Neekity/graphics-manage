<template>
  <v-app>
    <v-card  width="960px" class="mx-auto my-2">
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 sm="6"
                >
            <v-text-field
                v-model="channelName"
                label="频道名称"
                :rules="[value => !!value || '请输入频道名称！',]"
                dense
                outlined
            ></v-text-field>
          </v-col>
        </v-row>
      </v-card-title>
      <v-card-text class="text-left">
        <v-row>
          <v-col
              cols="12"
              sm="6"
          >
            <v-select
                v-model="owners"
                :items="allUsers"
                chips
                label="频道创作者"
                multiple
                outlined
                dense
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
                :items="items"
                selectable
                return-object
                dense
                hoverable
                :open.sync="open"
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
                    mdi-account
                  </v-icon>
                  {{ node.name }}
                </v-chip>
              </v-scroll-x-transition>
            </v-card-text>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-btn
            color="success"
            @click="store"
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
  name: "EditChannel",
  data() {
    return {
      open: [],
      owners: [],
      allUsers: [],
      channelId: parseInt(this.$route.query.id),
      allOpened: false,
      disableChannel: false,
      lastOpen: [],
      overlay: 0,
      channelName: '',
      search: null,
      selection: [],
      items: [{id: 0, name: '全选', children: []}],
    }
  },
  created() {
    this.getAllAudience();
  },
  methods: {
    getChannelAudience() {
      this.overlay += 1;
      this.$axios.post('/backend/channel/detail', {channel_id: this.channelId})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.selection = resData.data.audiences;
              this.channelName = resData.data.name;
              this.owners = resData.data.owners;
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
          .finally(() => this.overlay -= 1);
    },
    getAllAudience() {
      this.overlay += 1;
      this.$axios.post('/backend/user',{name: ''})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.items[0].children = resData.data;
              const that=this
              resData.data.forEach(function (element) {
                that.allUsers.push(element.name)
              });
              if (this.channel) {
                this.getChannelAudience()
              }
            } else {
              this.$toast('获取所有用户出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取所有用户出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    store() {
      this.overlay += 1;
      let receivers = [];
      this.selection.forEach(function (element) {
        receivers.push({id: element.id, name: element.name})
      });
      this.$axios.post('/backend/channel/store', {
        id: this.channelId,
        name: this.channelName,
        owners: this.owners,
        audiences: receivers
      })
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.$toast("保存成功！", {
                type: 'success',
                timeout: 2000
              });
              location.href = '/channel';
            } else {
              this.$toast('保存频道出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('保存频道出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
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
