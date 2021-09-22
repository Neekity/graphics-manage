<template>
  <div>
    <v-hover v-slot="{ hover }">
      <v-card
          :elevation="hover ? 12 : 0"
          :class="{ 'on-hover': hover }"
      >
        <v-img
            height="135px"
            :src="selectImage.url"
            aspect-ratio="2.347"
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
                    @click="dialog=true;stepper=2"
                >
                  <v-icon
                      :class="{ 'show-btns': hover }"
                      :color="transparent"
                  >mdi-pencil
                  </v-icon>
                </v-btn>
                <v-btn
                    :class="{ 'blue darken-2 mx-2': hover }"
                    :color="transparent"
                    icon
                    @click="dialog=true;stepper=1"
                >
                  <v-icon
                      :class="{ 'show-btns': hover }"
                      :color="transparent"
                  >mdi-swap-horizontal
                  </v-icon>
                </v-btn>
              </div>
            </v-row>
          </v-card-actions>
        </v-img>
      </v-card>
    </v-hover>
    <v-dialog v-model="dialog" width="960">
      <v-stepper
          v-model="stepper"
          vertical
      >
        <v-stepper-step
            :complete="stepper > 1"
            step="1"
        >
          选择封面
        </v-stepper-step>

        <v-stepper-content step="1">
          <v-card>
            <v-card-text>
              <v-card>
                <v-card-title>
                  <v-row class="p-2">
                    <v-spacer></v-spacer>
                    <v-tooltip bottom>
                      <template v-slot:activator="{ on, attrs }">
                        <v-btn
                            v-bind="attrs"
                            v-on="on"
                            icon
                            color="blue darken-2"

                        >
                          <v-icon>mdi-image-plus</v-icon>
                        </v-btn>
                      </template>
                      <span>上传图片</span>
                    </v-tooltip>

                  </v-row>
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
                                :color="selectImage.id===image.id?'green lighten-4':undefined"
                                @click="selectImage=image"
                            >
                              <v-img
                                  height="250"
                                  contain
                                  :src="image.url"
                              >
                              </v-img>
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
                          :src="preview_image_url"
                      >
                      </v-img>
                    </v-card>
                  </v-dialog>
                </v-card-text>
              </v-card>
            </v-card-text>
            <v-divider class="mx-4"></v-divider>
            <v-card-actions>
              <v-btn
                  color="primary"
                  id="init_editor"
                  @click="stepper = 2"
              >
                下一步
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-stepper-content>
        <v-stepper-step
            :complete="stepper > 2"
            step="2"
        >
          编辑封面
        </v-stepper-step>

        <v-stepper-content step="2">
          <v-card>
            <v-card-title>
              <v-radio-group
                  v-model="row"
                  row
              >
                <div class="px-2"><strong>请选取裁剪比例</strong></div>
                <v-radio
                    label="2.35:1"
                    value="radio-1"
                    @click="option.fixedNumber=[2.35,1]"
                ></v-radio>
                <v-radio
                    label="1:1"
                    value="radio-2"
                    @click="option.fixedNumber=[1,1]"
                ></v-radio>
              </v-radio-group>
            </v-card-title>
            <v-card-text>
              <v-row>
                <v-col>
                  <div style="width: 600px;height: 400px">
                    <vue-cropper ref="cropper" :img="selectImage.url" :output-size="option.size"
                                 :output-type="option.outputType"
                                 :info="true" :fixed="option.fixed" :fixed-number="option.fixedNumber"
                                 :can-move="option.canMove" :can-move-box="option.canMoveBox"
                                 :fixed-box="option.fixedBox"
                                 :original="option.original" :high="option.high"
                                 :auto-crop="option.autoCrop" :auto-crop-width="option.autoCropWidth"
                                 :auto-crop-height="option.autoCropHeight" :center-box="option.centerBox"
                                 mode="cover" :max-img-size="option.max" @real-time="realTime"></vue-cropper>
                  </div>
                </v-col>
                <v-col>
                  <v-row v-show="row==='radio-1'">
                    <h2 class="pa-4">2.35:1预览效果</h2>
                    <v-img
                        width="220px"
                        :src="preview_url_1"
                        aspect-ratio="2.35"
                    ></v-img>
                  </v-row>
                  <v-row v-show="row==='radio-2'">
                    <h2 class="pa-4">1:1预览效果</h2>
                    <v-img
                        width="220px"
                        :src="preview_url_2"
                        aspect-ratio="1"
                    ></v-img>
                  </v-row>
                </v-col>
              </v-row>
            </v-card-text>
            <v-card-actions>
              <v-btn
                  color="success"
                  @click="generate"
              >
                完成
              </v-btn>
              <v-btn
                  color="text"
                  @click="stepper=1"
              >
                上一步
              </v-btn>
              <v-btn
                  color="text"
                  @click="dialog=false"
              >
                取消
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-stepper-content>
      </v-stepper>
    </v-dialog>
  </div>
</template>

<script>
export default {
  name: "CoverImage",
  props: {
    parentImage: Object,
    channel: Number,
  },
  created() {
    this.selectImage = this.parentImage ?? {
      id: 0,
      name: '',
      url: '',
    }
  },
  data() {
    return {
      isImageActive: null,
      selectImage: {},
      images: [],
      stepper: 1,
      row: 'radio-1',
      preview_image_url: "",
      dialog: false,
      preview_url_1: "",
      preview_url_2: "",
      transparent: 'rgba(255, 255, 255, 0)',
      option: {
        fixedNumber: [2.34, 1],
        size: 1,
        fixed: true,
        full: false,
        outputType: 'png',
        canMove: true,
        fixedBox: false,
        original: false,
        canMoveBox: true,
        autoCrop: true,
        // 只有自动截图开启 宽度高度才生效
        autoCropWidth: 400,
        autoCropHeight: 400,
        centerBox: true,
        high: true,
        max: 99999
      },
    }
  },
  methods: {
    getImages() {
      this.overlay = true;
      this.$http.post('/graphics/material/list', {type: 'image', channel_id: this.channel})
          .then(response => {
            console.log(JSON.stringify(response.data));
            let resData = response.data;
            if (resData.code === 0) {
              this.images = resData.data;
            } else {
              this.$toast('获取图片出错：' + resData.message, {
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
    generate() {
      this.$refs.cropper.getCropBlob((data) => {
        console.log(data);
        var img = window.URL.createObjectURL(data)
        this.$emit('getCoverImageUrl', img)
        this.dialog = false;
      })

    },
    realTime() {
      this.$refs.cropper.getCropData((data) => {
        if (this.row === 'radio-1') {
          this.preview_url_1 = data
        } else {
          this.preview_url_2 = data
        }
      })
    }
  }
}
</script>

<style scoped>
.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
</style>