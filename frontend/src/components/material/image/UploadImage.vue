<template>
  <v-app>
    <v-card width="660px" class="mx-auto my-12">
      <v-card-title>
        <div class="text-h5">上传图片</div>
      </v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="12" lg="6"
                 md="6"
                 sm="12">
            <v-select
              v-model="channel"
              item-text="name"
              item-value="id"
              :items="channelItems"
              label="图片所属频道"
              :rules="[value => !!value || '请选择上传的频道！',]"
              outlined
            ></v-select>
          </v-col>
          <v-col cols="12" lg="6"
                 md="6"
                 sm="12"
          >
            <v-file-input
              label="点击选择图片"
              filled
              dense
              accept="image/gif,image/jpeg,image/jpg,image/png,image/svg"
              v-model="imageFile"
              prepend-icon="mdi-camera"
            ></v-file-input>
          </v-col>
        </v-row>
      </v-card-text>
     <v-card-actions>
         <v-spacer></v-spacer>
         <v-btn
           color="success"
           :disabled="!imageFile || !channel"
           @click="upload"
         >
           确认上传
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
  name: "UploadImage",
  data() {
    return {
      imageFile: null,
      channel: null,
      channelItems: [],
      overlay: false,
    }
  },
  created() {
    this.getChannel();
  },
  methods: {
    upload() {
      this.uploadImageDialog = false;
      this.overlay = true;
      let formData = new FormData();
      formData.append('graphics_img', this.imageFile, this.imageFile.name);
      formData.append('channel_id', this.channel);
      this.$http.post('/graphics/material/owner/image/upload', formData)
        .then(response => {
          console.log(JSON.stringify(response.data));
          let resData = response.data;
          if (resData.code === 0) {
            this.$toast("上传成功！", {
              type: 'success',
              timeout: 2000
            });
            location.href = '/graphics/material/owner/image';
          } else if (resData.code === 40 || resData.code === 1000004001) {
            this.$toast('上传出错：' + resData.message + '【' +
              resData.data.errors[Object.keys(resData.data.errors)[0]][0] + '】', {
              type: 'error',
            });
          } else {
            this.$toast('上传出错：' + resData.message, {
              type: 'error',
            });
          }
        })
        .catch(error => {
          console.log(error);
          this.$toast('上传出错：服务器出错！', {
            type: 'error',
            timeout: 2000,
          });
        })
        .finally(() => this.overlay = false);
    },
    getChannel() {
      this.overlay = true;
      this.$http.post('/graphics/material/owner/channel')
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
  }
}
</script>

<style scoped>

</style>
