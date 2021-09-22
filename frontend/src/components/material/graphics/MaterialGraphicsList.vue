<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row class="pl-2">
          <v-col
              cols="12"
              lg="3"
              md="4"
              sm="6"
              xs="8"
          >
            <v-text-field
                v-model="materialName"
                label="请输入名称搜索"
                single-line
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
          <v-spacer></v-spacer>
          <v-col
              cols="12"
              lg="1"
              md="2"
              sm="2"
              xs="4"
          >
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    v-bind="attrs"
                    v-on="on"
                    icon
                    color="blue darken-2"
                    href="/graphics/material/graphics/create"
                >
                  <v-icon>mdi-plus</v-icon>
                </v-btn>
              </template>
              <span>新建图文模版</span>
            </v-tooltip>
          </v-col>
        </v-row>

      </v-card-title>
      <v-card-text>
        <v-lazy
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
                  <v-card-title class="text-h6">
                    {{ item.name }}
                  </v-card-title>
                  <div class="tinymce-viewer px-2">
                    <tiny-mce-viewer v-model="item.content" :skin-url="skinUrl"></tiny-mce-viewer>
                  </div>
                  <v-divider></v-divider>
                  <v-card-actions class="pt-0">
                    <div>更新于 {{ item.update_at }}</div>
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
                              @click="deleteTemplate()"
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
import TinyMceViewer from "../../tinymce/TinyMceViewer";

export default {
  name: "MaterialGraphicsList",
  components: {TinyMceViewer},
  props: {
    channel: Number,
    skinUrl: String,
  },
  data() {
    return {
      isActive: false,
      graphics: [],
      materialName: "",
      overlay: false,
      dialog: false,
      deleteId: 0,
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
    // this.getGraphics()
  },
  methods: {
    getGraphics() {
      this.overlay = true;
      this.$http.post('/graphics/material/list', {type: 'graphics', channel_id: this.channel, name: this.materialName})
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
      location.href = "/graphics/material/graphics/" + id + "/edit"
    },
    confirm(id) {
      this.dialog = true;
      this.deleteId = id
    },
    deleteTemplate() {
      this.dialog = false;
      this.overlay = true;
      this.$http.post('/graphics/material/delete', {id: this.deleteId})
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