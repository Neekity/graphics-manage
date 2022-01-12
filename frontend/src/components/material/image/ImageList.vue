<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 lg="3"
                 md="4"
                 sm="5"
                 xs="12"
          >
            <v-text-field
                v-model="materialName"
                label="请输入名称搜索"
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getImages"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-col cols="12"
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
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                  v-bind="attrs"
                  v-on="on"
                  icon
                  color="blue darken-2"
                  @click="uploadImageDialog=true"
              >
                <v-icon>mdi-image-plus</v-icon>
              </v-btn>
            </template>
            <span>上传</span>
          </v-tooltip>
        </v-row>
        <v-dialog
            v-model="uploadImageDialog"
            width="400"
        >
          <v-card>
            <v-card-title class="pa-4">
            </v-card-title>
            <v-card-text>
              <v-file-input
                  label="点击选择图片"
                  filled
                  accept="image/png, image/jpeg, image/bmp"
                  v-model="imageFile"
                  prepend-icon="mdi-camera"
              ></v-file-input>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                  color="success"
                  @click="upload"
              >
                确认上传
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-card-title>
      <v-card-text>
        <v-lazy
            v-model="isImageActive"
            :options="{
          threshold: .5
        }"
            min-height="200"
            transition="fade-transition"
        >
          <v-container fluid class="text-center">
            <v-row align="center"
                   justify="center"
                   class="fill-height"
            >
              <v-col
                  v-for="image in images"
                  :key="image.id"
                  cols="12"
                  lg="3"
                  md="4"
                  sm="6"
                  xs="12"
              >
                <v-hover v-slot="{ hover }">
                  <v-card
                      :elevation="hover ? 12 : 0"
                      :class="{ 'on-hover': hover }"
                  >
                    <v-img
                        height="250"
                        contain
                        :src="image.url"
                    >
                      <v-card-actions>
                        <v-row
                            class="fill-height flex-column"
                            justify="space-between"
                        >
                          <div class="align-self-center pt-2">
                            <v-btn
                                :class="{ 'blue darken-2 mx-2': hover }"
                                :color="transparent"
                                icon
                                @click="preview(image.url)"
                            >
                              <v-icon
                                  :class="{ 'show-btns': hover }"
                                  :color="transparent"
                              >mdi-eye
                              </v-icon>
                            </v-btn>

                            <v-btn
                                :class="{ 'red darken-2 mx-2': hover }"
                                :color="transparent"
                                icon
                                @click="deleteImage(image.id)"
                            >
                              <v-icon
                                  :class="{ 'show-btns': hover }"
                                  :color="transparent"
                              >mdi-delete
                              </v-icon>
                            </v-btn>
                          </div>
                        </v-row>
                      </v-card-actions>
                    </v-img>
                    <div class="px-4 py-2 subheading text-truncate">
                      {{ image.url }}
                    </div>
                  </v-card>
                </v-hover>
              </v-col>
            </v-row>
          </v-container>
        </v-lazy>
        <v-dialog
            v-model="dialog"
        >
          <v-card>
            <v-img
                :src="previewImageUrl"
            >
            </v-img>
          </v-card>
        </v-dialog>
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
export default {
  name: "ImageList",
  data() {
    return {
      channel: 0,
      uploadImageDialog: false,
      overlay: 0,
      dialog: false,
      imageFile: null,
      previewImageUrl: "",
      isImageActive: null,
      images: [],
      transparent: 'rgba(255, 255, 255, 0)',
      materialName: "",
      channelItems: [],

    }
  },
  watch: {
    channel: {
      handler() {
        this.getImages()
      },
    },
  },
  created() {
    this.getChannel();
  },
  mounted() {
    this.getImages()
  },
  methods: {
    getChannel() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/channels')
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
    upload() {
      this.uploadImageDialog = false;
      this.overlay += 1;
      let formData = new FormData();
      formData.append('graphics_img', this.imageFile, this.imageFile.name);
      formData.append('channel_id', this.channel);
      this.$graphicsHttp('post','/material/image/upload', formData)
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.$toast("上传成功！", {
                type: 'success',
                timeout: 2000
              });
              this.getImages()
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
            this.$toast('保存出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    preview(url) {
      this.previewImageUrl = url;
      this.dialog = true;
    },
    deleteImage(id) {
      this.overlay += 1
      this.$graphicsHttp('post', '/material/delete', {id: id})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.$toast('删除图片成功！' + resData.message, {
                type: 'success',
                timeout: 2000,
              });
              this.getImages();
            } else {
              this.$toast('删除图片出错：' + resData.message, {
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
          .finally(() => this.overlay -= 1);
    },
    getImages() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/list', {
        type: 'image',
        channel_id: this.channel,
        name: this.materialName
      }).then(response => {
        let resData = response.data;
        if (resData.code === 0) {
          this.images = resData.data || [];
        } else {
          this.$toast('获取图片出错：' + resData.message, {
            type: 'error',
          });
        }
      }).catch(error => {
        console.log(error);
        this.$toast('获取图片出错：服务器出错！', {
          type: 'error',
          timeout: 2000,
        });
      }).finally(() => this.overlay -= 1);

    },
  },
}
</script>

<style scoped>
.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
</style>