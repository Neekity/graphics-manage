<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 lg="3"
                 md="4"
                 sm="6"
                 xs="12"
          >
            <v-text-field
                v-model="materialName"
                label="请输入标题"
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getGraphics"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-col cols="12"
                 lg="3"
                 md="4"
                 sm="6"
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
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                  v-bind="attrs"
                  v-on="on"
                  icon
                  color="blue darken-2"
                  href="/material/create"
              >
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </template>
            <span>新建图文素材</span>
          </v-tooltip>
        </v-row>
      </v-card-title>
      <v-card-text>
        <v-data-table
            v-if="showTable"
            :headers="headers"
            :items="graphics"
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
            <v-icon
                small
                class="mr-2"
                color="blue darken-2"
                @click="edit(item.id)"
            >
              mdi-pencil
            </v-icon>
            <v-icon
                small
                color="red darken-2"
                @click="confirm(item.id)"
            >
              mdi-delete
            </v-icon>
          </template>
          <template v-slot:no-data>
            <v-btn
                color="primary"
                @click="initialize"
            >
              Reset
            </v-btn>
          </template>
          <template v-slot:item.name="{ item }">
            <TextTruncate :text="item.name"></TextTruncate>
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
                  v-for="item in graphics"
                  :key="item.id"
                  cols="12"
                  lg="3"
                  md="4"
                  sm="6"
                  xs="12"
              >
                <v-card>
                  <v-img
                      :src="item.url"
                      aspect-ratio="2.347"
                  >
                  </v-img>
                  <v-card-title class="text-h6">
                    <TextTruncate :text="item.name" max-width="200px"></TextTruncate>
                  </v-card-title>
                  <v-card-subtitle class="text-left">
                    <TextTruncate :text="item.abstract?item.abstract:'<未填写摘要>'" max-width="200px"></TextTruncate>
                  </v-card-subtitle>
                  <v-divider></v-divider>
                  <v-card-actions class="pt-0">
                    <div>更新于 {{ item.updated_at }}</div>
                    <v-spacer></v-spacer>
                    <v-btn icon color="blue darken-2" @click="edit(item.id)">
                      <v-icon>mdi-pencil</v-icon>
                    </v-btn>
                    <v-btn icon color="blue darken-2" @click="confirm(item.id)">
                      <v-icon>mdi-delete</v-icon>
                    </v-btn>
                    <v-dialog
                        v-model="dialog"
                        width="500"
                    >
                      <v-card>
                        <v-card-title class="text-h5 light-red lighten-2">
                          确认删除？
                        </v-card-title>
                        <v-card-actions>
                          <v-spacer></v-spacer>
                          <v-btn @click="dialog = false" text>
                            取消
                          </v-btn>
                          <v-btn
                              color="success"
                              text
                              @click="deleteGraphics()"
                          >
                            确认
                          </v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-dialog>
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
import TextTruncate from "../../TextTruncate";

export default {
  name: "MaterialGraphicsList",
  components: {TextTruncate},
  data() {
    return {
      channel: 0,
      isActive: false,
      showTable: false,
      graphics: [],
      materialName: "",
      overlay: false,
      dialog: false,
      deleteId: 0,
      channelItems:[],
      headers: [{
        text: '标题',
        align: 'start',
        sortable: false,
        value: 'name',
      },
        {text: '摘要', value: 'abstract'},
        {text: '频道', value: 'channel_id'},
        {text: '更新时间', value: 'updated_at'},
        {text: 'Actions', value: 'actions', sortable: false},
      ]
    }
  },
  watch: {
    channel: {
      handler() {
        this.getGraphics()
      }
    },
  },
  mounted() {
    this.getGraphics()
    this.getChannel()
  },
  methods: {
    getChannel() {
      this.overlay = true;
      this.$graphicsHttp('post', 'material/channels')
          .then(response => {
            console.log(JSON.stringify(response.data));
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
    getGraphics() {
      this.overlay = true;
      this.$graphicsHttp('post', 'material/list', {type: 'graphics', name: this.materialName,channel_id:this.channel})
          .then(response => {
            console.log(JSON.stringify(response.data));
            let resData = response.data;
            if (resData.code === 0) {
              this.graphics = resData.data;
            } else {
              this.$toast('获取图文素材出错：' + resData.message, {
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
    edit(id) {
      location.href = '/material/graphics/edit?id='+id;
    },
    confirm(id) {
      this.dialog = true;
      this.deleteId = id
    },
    deleteGraphics() {
      this.dialog = false;
      this.overlay = true;
      this.$axios.post('material/delete', {id: this.deleteId})
          .then(response => {
            console.log(JSON.stringify(response.data));
            let resData = response.data;
            if (resData.code === 0) {
              this.$toast("删除成功！", {
                type: 'success',
                timeout: 2000
              });
              this.getGraphics()
            } else {
              this.$toast('删除数据出错：' + resData.message, {
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
<style>
.tinymce-viewer {
  height: 150px !important;
  overflow: hidden;
}

.tinymce-viewer img {
  height: auto;
  width: 100%;
}
</style>