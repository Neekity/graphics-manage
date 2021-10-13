<template>
  <v-app>
    <v-card class="my-2">
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
          <v-col>
            <v-text-field
                v-model="searchAudience"
                label="输入关键字搜索听众"
                clearable
                dense
                outlined
                @input="handleSearchAudience"
                clear-icon="mdi-close-circle-outline"
            ></v-text-field>
            <v-treeview
                ref="audienceItems"
                v-model="audienceSelection"
                :items="audienceItems"
                selectable
                return-object
                dense
                hoverable
                :open.sync="audienceOpen"
                :search="searchAudience"
            ></v-treeview>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col>
            <v-card-text>
              <div class="text-h5" >听众</div>
              <v-scroll-x-transition
                  group
                  hide-on-leave
              >
                <v-chip
                    v-for="(node, i) in audienceSelection"
                    :key="i"
                    color="primary"
                    outlined
                    small
                    close
                    class="ma-1"
                    @click:close="audienceSelection.splice(i,1)"
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
        <v-row>
          <v-col>
            <v-text-field
                v-model="searchChannelOwner"
                label="输入关键字搜索频道拥有者"
                clearable
                dense
                outlined
                @input="handleSearchChannelOwners"
                clear-icon="mdi-close-circle-outline"
            ></v-text-field>
            <v-treeview
                ref="channelOwnerItems"
                v-model="channelOwnerSelection"
                :items="channelOwnerItems"
                selectable
                return-object
                dense
                hoverable
                :open.sync="channelOwnerOpen"
                :search="searchChannelOwner"
            ></v-treeview>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col>
            <v-card-text>
              <div class="text-h5" >频道拥有者</div>
              <v-scroll-x-transition
                  group
                  hide-on-leave
              >
                <v-chip
                    v-for="(node, i) in channelOwnerSelection"
                    :key="i"
                    color="primary"
                    outlined
                    small
                    close
                    class="ma-1"
                    @click:close="channelOwnerSelection.splice(i,1)"
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
      channelOwnerOpen: [],
      audienceOpen: [],
      channelId: parseInt(this.$route.query.id),
      channelOwnerAllOpened: false,
      audienceAllOpened: false,
      disableChannel: false,
      audienceLastOpen: [],
      channelOwnerLastOpen: [],
      overlay: 0,
      channelName: '',
      searchChannelOwner: null,
      searchAudience: null,
      channelOwnerSelection: [],
      audienceSelection: [],
      audienceItems: [{id: 0, name: '全选', children: []}],
      channelOwnerItems: [{id: 0, name: '全选', children: []}],
    }
  },
  created() {
    this.getAllAudience();
  },
  methods: {
    getChannelDetail() {
      this.overlay += 1;
      this.$graphicsHttp('post','/channel/detail', {id: this.channelId})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.audienceSelection = JSON.parse(resData.data.audiences);
              this.channelOwnerSelection = JSON.parse(resData.data.channel_owners);
              this.channelName = resData.data.name;
            } else {
              this.$toast('获取频道信息出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取频道信息出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    getAllAudience() {
      this.overlay += 1;
      this.$graphicsHttp('post','/user',{name: ''})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.audienceItems[0].children = resData.data;
              this.channelOwnerItems[0].children = resData.data;
              if (this.channelId) {
                this.getChannelDetail()
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
      let audiences = [];
      let owners = [];
      this.audienceSelection.forEach(function (element) {
        audiences.push({id: element.id, name: element.name})
      });
      this.channelOwnerSelection.forEach(function (element) {
        owners.push({id: element.id, name: element.name})
      });
      this.$graphicsHttp('post','/channel/store', {
        id: this.channelId,
        name: this.channelName,
        owners: owners,
        audiences: audiences
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
    handleSearchAudience: function (val) {
      if (val) {
        if (!this.audienceAllOpened) {
          this.audienceLastOpen = this.open;
          this.audienceAllOpened = true;
          this.$refs.audienceItems.updateAll(true);
        }
      } else {
        this.$refs.audienceItems.updateAll(false);
        this.audienceAllOpened = false;
        this.audienceOpen = this.audienceLastOpen;
      }
    },
    handleSearchChannelOwners: function (val) {
      if (val) {
        if (!this.channelOwnerAllOpened) {
          this.channelOwnerLastOpen = this.open;
          this.channelOwnerAllOpened = true;
          this.$refs.channelOwnerItems.updateAll(true);
        }
      } else {
        this.$refs.channelOwnerItems.updateAll(false);
        this.channelOwnerAllOpened = false;
        this.channelOwnerOpen = this.channelOwnerLastOpen;
      }
    }
  },
}
</script>
