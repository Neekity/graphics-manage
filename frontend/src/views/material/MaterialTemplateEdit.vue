<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 lg="2"
                 md="3"
                 sm="4"
                 xs="6">
            <v-select
                v-model="channel"
                item-text="name"
                item-value="id"
                :items="channelItems"
                label="所属频道"
                dense
                outlined
                disabled
            ></v-select>
          </v-col>
        </v-row>
      </v-card-title>
      <v-card-text>
        <v-card width="578px" class="mx-auto my-12">
          <v-card-title>
            <v-text-field id="title" v-model="materialName" outlined dense label="名称"></v-text-field>
          </v-card-title>
          <v-card-text>
            <v-row align="center" class="mx-0 tinymce-content">
              <tinymce-editor v-model="content" base-url="tinymce/" :channel="channel"></tinymce-editor>
            </v-row>
          </v-card-text>
          <v-divider class="mx-4"></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
                color="success"
                @click="store"
            >
              保存模版
            </v-btn>
          </v-card-actions>
        </v-card>
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
import TinymceEditor from "../../components/tinymce/TinymceEditor";

export default {
  name: "MaterialTemplateEdit",
  components: {TinymceEditor},
  props: ['skinUrl', 'materialId', 'channel'],
  data() {
    return {
      materialName: "",
      content: '',
      overlay:false,
      channelItems:[],
    }
  },
  mounted() {

  },
  methods: {
    getChannel() {
      this.overlay = true;
      this.$http.post('/graphics/material/channel')
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
    getMaterial() {
      if (this.materialId === 0) {
        return;
      }
      this.overlay = true;
      this.$http.post('/graphics/material/data', {id: this.materialId})
          .then(response => {
            console.log(JSON.stringify(response.data));
            let resData = response.data;
            if (resData.code === 0) {
              this.content = resData.data.content;
              this.materialName = resData.data.name;
            } else {
              this.$toast('获取模版出错：' + resData.message, {
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
    store() {
      this.overlay = true;
      let url = '/graphics/material/update';
      if (this.materialId == 0) {
        url = '/graphics/material/store';
      }
      this.$http.post(url, {
        name: this.materialName,
        content: this.content,
        id: this.materialId,
        channel_id: this.channel,
        type: 'template',
      }).then(response => {
        console.log(JSON.stringify(response.data));
        let resData = response.data;
        if (resData.code === 0) {
          this.$toast("保存成功！", {
            type: 'success',
            timeout: 2000
          });
        } else if (resData.code === 40 || resData.code === 1000004001) {
          this.$toast('保存出错：' + resData.message + '【' +
              resData.data.errors[Object.keys(resData.data.errors)[0]][0] + '】', {
            type: 'error',
          });
        } else {
          this.$toast('保存出错：' + resData.message, {
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
img {
  display: block;
  width: 550px;
  height: auto;
}
</style>