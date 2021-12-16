<template>
  <v-app>
    <v-card width="660px" class="mx-auto my-12">
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 lg="9"
                 md="8"
                 sm="6"
                 xs="12">
            <v-text-field v-model="materialName" outlined dense label="模版名称"
                          :rules="[value => !!value || '请输入模版名称！',]"></v-text-field>
          </v-col>
          <v-col cols="12"
                 lg="3"
                 md="4"
                 sm="6"
                 xs="12">
            <v-select
                v-model="channel"
                item-text="name"
                item-value="id"
                :items="channelItems"
                label="素材所属频道"
                :disabled="channelDisabled"
                dense
                outlined
            ></v-select>
          </v-col>
        </v-row>
      </v-card-title>
      <v-card-text>
        <v-row align="center" class="mx-0 tinymce-content">
          <tinymce-editor v-model="content" :base-url="baseUrl" :channel="channel" :id="'tinymce-template-'+materialId"></tinymce-editor>
        </v-row>
      </v-card-text>
      <v-divider class="mx-4"></v-divider>
      <v-card-subtitle>
        <v-row class="pt-3">
          <v-col cols="12"
                 lg="6"
                 md="6"
                 sm="12"
                 xs="12">
            <v-textarea label="模版描述" v-model="abstract" filled
                        rows="3" row-height="30"
            ></v-textarea>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
      </v-card-subtitle>
      <v-divider class="mx-4"></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="success"
            :disabled="!materialName || !channel"
            @click="store"
        >
          保存模版
        </v-btn>
        <v-btn
            color="text"
            href="/graphics/material/owner/template"
        >
          返回
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
import TinymceEditor from "../../components/tinymce/TinymceEditor";

export default {
  name: "MaterialTemplateEdit",
  components: {TinymceEditor},
  props: ['baseUrl', 'parentMaterialId'],
  data() {
    const _this = this;
    return {
      channel: null,
      materialName: "",
      content: '',
      abstract: "",
      overlay: false,
      channelItems: [],
      materialId: this.parentMaterialId,
      channelDisabled: false,
    }
  },
  created() {
    this.getChannel();
    this.getMaterial();
  },
  methods: {
    getChannel() {
      this.overlay = true;
      this.$graphicsHttp('post','/graphics/material/owner/channel')
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
      if (this.materialId == 0) {
        return;
      }
      this.overlay = true;
      this.$http.post('/graphics/material/owner/data', {id: this.materialId})
          .then(response => {
            console.log(JSON.stringify(response.data));
            let resData = response.data;
            if (resData.code === 0) {
              this.content = resData.data.content;
              this.materialName = resData.data.name;
              this.abstract = resData.data.abstract;
              this.channel = resData.data.channel_id;
              this.channelDisabled = true;
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
      let url = '/graphics/material/owner/update';
      if (this.materialId == 0) {
        url = '/graphics/material/owner/store';
      }
      this.$http.post(url, {
        name: this.materialName,
        content: this.content,
        id: this.materialId,
        abstract: this.abstract,
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
          this.materialId = resData.data.id;
          this.channelDisabled = true;
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
.tox-tinymce {
  border-bottom:none !important;
  border-left:none !important;
  border-right:none !important;
}
</style>
