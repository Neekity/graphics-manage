<template>
  <v-app>
    <v-card width="660px" class="mx-auto my-12">
      <v-card-title>
        <v-row>
          <v-col cols="12">
            <v-text-field id="title" v-model="materialName" outlined dense label="标题" hide-details="auto"
                          :rules="[value => !!value || '请输入标题！',]"></v-text-field>
          </v-col>
          <v-col cols="12"
                 lg="6"
                 md="6"
                 sm="6"
                 xs="12">
            <v-text-field v-model="author" outlined dense label="作者（非必填）"></v-text-field>
          </v-col>
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
                label="素材所属频道"
                :rules="[value => !!value || '请选择素材想要上传的频道！',]"
                :disabled="channelDisabled"
                dense
                outlined
            ></v-select>
          </v-col>
        </v-row>
      </v-card-title>
      <v-card-text>
        <v-row
            align="center"
            class="mx-0 tinymce-content"
        >
          <div v-if="!showEditor">
            <v-alert
                dense
                outlined
                type="info"
            >选择频道后方可编辑正文内容
            </v-alert>
          </div>
          <tinymce-message-editor @hook:destroyed="editorInit" v-if="showEditor" v-model="content"
                                  :editor-templates="channelTemplates" @getImages="getImages"
                                  @showSelectMaterialModal="showSelectMaterialModal"
                                  :id="'tinymce-graphics-'+materialId"
                                  :base-url="baseUrl" :channel="channel"></tinymce-message-editor>
        </v-row>
      </v-card-text>
      <v-divider class="mx-4"></v-divider>
      <v-card-subtitle>
        <v-row>
          <div class="text-h5 pa-3">封面和摘要</div>
          <v-spacer></v-spacer>
        </v-row>

        <v-row>
          <v-col>
            <v-hover v-slot="{ hover }">
              <v-card
                  :elevation="hover ? 12 : 0"
                  :class="{ 'on-hover': hover }"
              >
                <v-img
                    height="135px"
                    :src="imageUrl"
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
                            @click="selectDialog=true"
                        >
                          <v-icon
                              :class="{ 'show-btns': hover }"
                              :color="transparent"
                          >mdi-swap-horizontal
                          </v-icon>
                        </v-btn>
                        <v-btn
                            :class="{ 'red darken-2 mx-2': hover }"
                            :color="transparent"
                            icon
                            @click="imageUrl=''"
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
              </v-card>
            </v-hover>
            <v-dialog v-model="selectDialog" width="960">
              <v-card>
                <v-card-title>
                  <v-row>
                    <div class="text-h5 pa-3">选择封面图（建议尺寸940*400）</div>
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
                              accept="image/gif,image/jpeg,image/jpg,image/png,image/svg"
                              v-model="imageFile"
                              prepend-icon="mdi-camera"
                          ></v-file-input>
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
                    </v-dialog>
                    <v-overlay :value="overlay">
                      <v-progress-circular
                          indeterminate
                          color="primary"
                      ></v-progress-circular>
                    </v-overlay>
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
                                @click="selectImage=image;imageUrl=image.url;selectDialog=false"
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
                </v-card-text>
              </v-card>
            </v-dialog>
            <v-dialog v-model="insertMaterialDialog" width="960">
              <v-card>
                <v-card-text>
                  <v-tabs v-model="tab">
                    <v-tabs-slider></v-tabs-slider>
                    <v-tab
                        v-for="item in tabItems"
                        :key="item"
                    >{{ item }}
                    </v-tab>
                  </v-tabs>
                  <v-tabs-items v-model="tab">
                    <v-tab-item
                        key="图片">
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
                                    @click="insertImageToContent(image.url)"
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
                    </v-tab-item>
                    <v-tab-item
                        key="模版">
                      <v-lazy
                          v-model="isTemplateActive"
                          :options="{
          threshold: .5
        }"
                          min-height="200"
                          transition="fade-transition"
                      >
                        <v-container fluid>
                          <v-row>
                            <v-col
                                v-for="template in templates"
                                :key="template.id"
                                cols="12"
                                lg="3"
                                md="4"
                                sm="6"
                                xs="12"
                            >
                              <v-card>
                                <v-card-title class="text-h6">
                                  {{ template.name }}
                                </v-card-title>
                                <div class="tinymce-viewer px-2">
                                  <tinymce-viewer v-model="template.content" :base-url="baseUrl"></tinymce-viewer>
                                </div>
                                <v-divider></v-divider>
                                <v-card-actions class="pt-0">
                                  <div>更新于 {{ template.updated_at }}</div>
                                  <v-spacer></v-spacer>
                                  <v-tooltip bottom>
                                    <template v-slot:activator="{ on, attrs }">
                                      <v-btn
                                          v-bind="attrs"
                                          v-on="on"
                                          icon color="blue darken-2"  @click="insertTemplateToContent(template.content)"
                                      >
                                        <v-icon>mdi-tray-arrow-down</v-icon>
                                      </v-btn>
                                    </template>
                                    <span>点击插入模版</span>
                                  </v-tooltip>
                                </v-card-actions>
                              </v-card>
                            </v-col>
                          </v-row>
                        </v-container>
                      </v-lazy>
                    </v-tab-item>
                  </v-tabs-items>

                </v-card-text>
              </v-card>
            </v-dialog>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col>
            <v-textarea label="摘要" v-model="abstract" filled
                        rows="4"
                        row-height="30"
            ></v-textarea>
          </v-col>
        </v-row>
      </v-card-subtitle>
      <v-divider class="mx-4"></v-divider>
      <v-card-actions class="pa-3">
        <v-spacer></v-spacer>
        <v-btn
            color="success"
            :disabled="!materialName || !channel"
            @click="store(false)"
        >
          保存
        </v-btn>
        <v-btn
            color="text"
            :disabled="!materialName || !channel"
            @click="storeAndPublish"
        >
          保存并群发
        </v-btn>
        <v-btn
            color="text"
            href="/graphics/material/owner/graphics"
        >
          返回
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-dialog v-model="publishDialog" width="960">
      <v-card class="overflow-y-auto" style="max-height: 600px" v-scroll.self="onScroll">
        <v-card-text>
          <v-row class="pt-6">
            <v-col cols="12">
              <div>
                <h5>选择听众</h5>
                <v-row>
                  <v-col>
                    <v-text-field
                        v-model="audiencesSearch"
                        label="输入关键字搜索听众"
                        clearable
                        dense
                        outlined
                        @input="handleAudiencesSearch"
                        clear-icon="mdi-close-circle-outline"
                    ></v-text-field>
                    <v-treeview
                        :key="audiencesTreeKey"
                        ref="audiencesTree"
                        v-model="audiencesTreeSelection"
                        selection-type="independent"
                        item-text="text"
                        :items="audiencesTree"
                        selectable
                        return-object
                        dense
                        hoverable
                        :open.sync="audiencesOpen"
                        :search="audiencesSearch"
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
                            v-for="(node, i) in audiencesSelection"
                            :key="i"
                            color="primary"
                            outlined
                            small
                            close
                            class="ma-1"
                            @click:close="removeAudience(i)"
                        >
                          <v-icon
                              left
                              small
                          >
                            mdi-account
                          </v-icon>
                          {{ node.text }}
                        </v-chip>
                      </v-scroll-x-transition>
                    </v-card-text>
                  </v-col>
                </v-row>
              </div>
            </v-col>
          </v-row>
          <v-row align="center" justify="center">
            <v-col cols="12" md="3" sm="6" xs="12">
              <v-switch
                  v-model="timedSending"
                  label="定时群发"
              ></v-switch>
            </v-col>
            <v-col v-if="timedSending" cols="12" md="3" sm="6" xs="12">
              <date-picker class="my-date-picker-svg" v-model="sendTime" :disabled-date="notBeforeToday" type="datetime"
                           value-type="format" format="YYYY-MM-DD HH:mm:ss"></date-picker>
            </v-col>
            <v-spacer></v-spacer>
          </v-row>
        </v-card-text>
        <v-card-actions>
          <v-btn
              color="success"
              @click="publish"
              class="text-center"
          >
            群发
          </v-btn>
          <v-btn
              color="success"
              @click="scrollTarget"
              class="text-center"
          >
            滚动
          </v-btn>
        </v-card-actions>
      </v-card>

    </v-dialog>
    <v-overlay :value="overlay">
      <v-progress-circular
          indeterminate
          color="primary"
      ></v-progress-circular>
    </v-overlay>
  </v-app>
</template>

<script>
import TinymceViewer from "../../components/tinymce/TinymceViewer";
let moment = require("moment");
import TinymceMessageEditor from "../../components/tinymce/TinymceMessageEditor";
import 'vue2-datepicker/index.css';

export default {
  name: "MaterialGraphicsEdit",
  components: {TinymceViewer, TinymceMessageEditor},
  data() {
    return {
      tab:null,
      tabItems: ['图片', '模版'],
      baseUrl: '/tinymce/',
      uploadImageDialog: false,
      imageFile: null,
      audiencesAllOpened: false,
      audiencesLastOpen: [],
      audiencesOpen: [],
      audiencesSearch: null,
      audiencesSelection: [],
      audiencesTreeKey: 0,
      audiencesTreeSelection: [],
      audiencesTree: [{id: '0', text: '全部听众', children: []}],
      overlay: 0,
      timedSending: false,
      authorSwitch: false,
      showEditor: false,
      author: "",
      channel: null,
      sendTime: moment().format('YYYY-MM-DD HH:mm:ss'),
      materialId: parseInt(this.$route.query.id),
      channelItems: [],
      list: [],
      stepper: 1,
      abstract: '',
      publishDialog: false,
      channelDisabled: false,
      materialName: '',
      imageUrl: '',
      content: '',
      isImageActive: null,
      isTemplateActive: null,
      channelTemplates: [],
      templates: [],
      selectImage: {},
      images: [],
      selectDialog: false,
      insertMaterialDialog: false,
      transparent: 'rgba(255, 255, 255, 0)',
      disableNodeFinished: false,
      systemSet: false,
      disableNodeArray: [],
      enableNodeMaps: {},
      scrollEle: null,
    }
  },
  watch: {
    channel: {
      handler() {
        if (this.channel) {
          this.getChannelTemplates();
          this.getImages()
          this.getChannelAudience();
        }
      },
    },
    audiencesTreeSelection: {
      handler(newV, oldV) {
        if (this.systemSet) {
          return;
        }

        this.systemSet = true;
        let newIds = []
        let oldIds = []
        newV.forEach(function (item) {
          newIds.push(item.id)
        })
        oldV.forEach(function (item) {
          oldIds.push(item.id)
        })

        const that = this;
        //增加的节点
        let addIds = newIds.filter(function (v) {
          return oldIds.indexOf(v) === -1
        });
        if (addIds.length > 0) {
          addIds.forEach(function (id) {
            newV.forEach(function (item) {
              if (item.id == id) {
                if (item.children) {
                  item.children.forEach(function (children) {
                    that.disableNodes(children)
                  })
                }
              }
            })
          })
          this.audiencesTreeSelection = this.audiencesTreeSelection.concat(this.disableNodeArray)
        }

        //删除的节点
        let removedIds = oldIds.filter(function (v) {
          return newIds.indexOf(v) === -1
        });
        if (removedIds.length > 0) {
          removedIds.forEach(function (id) {
            oldV.forEach(function (item) {
              if (item.id == id) {
                if (item.children) {
                  item.children.forEach(function (children) {
                    that.enableNodes(children)
                  })
                }
              }
            })
          })
          this.removeDisabledNode();
        }
        this.audiencesSelection = [];
        for (let i = 0; i < this.audiencesTreeSelection.length; i++) {
          if (!this.audiencesTreeSelection[i].disabled) {
            this.audiencesSelection.push(this.audiencesTreeSelection[i])
          }
        }
        this.$nextTick(function () {
          this.systemSet = false;
          this.disableNodeArray = [];
          this.enableNodeMaps = {};
        });
        this.audiencesTreeKey += 1;
        if (this.scrollEle) {
          this.$nextTick(() => {
            this.scrollEle.target.scrollTop = this.scrollEleTop;
          })
        }
      },
    }
  },
  created() {
    this.getChannel();
    this.getMaterial();
  },
  methods: {
    scrollTarget() {
      this.scrollEle.target.scrollTop = this.scrollEleTop;
    },
    onScroll(e) {
      if (this.scrollEle == null) {
        this.scrollEle = e;
      }
      this.scrollEleTop = e.target.scrollTop;

    },
    removeAudience(i) {
      this.systemSet = true;
      let treeIdx = 0
      for (; treeIdx < this.audiencesTreeSelection.length; treeIdx++) {
        if (this.audiencesTreeSelection[treeIdx].id == this.audiencesSelection[i].id) {
          if (this.audiencesTreeSelection[treeIdx].children) {
            for (let childrenIdx = 0; childrenIdx < this.audiencesTreeSelection[treeIdx].children.length; childrenIdx++) {
              this.enableNodes(this.audiencesTreeSelection[treeIdx].children[childrenIdx])
            }
          }
          break;
        }
      }
      this.audiencesTreeSelection.splice(treeIdx, 1)
      this.audiencesSelection.splice(i, 1)
      this.removeDisabledNode();
      this.$nextTick(function () {
        this.systemSet = false;
        this.enableNodeMaps = {};
      });
      this.audiencesTreeKey += 1;
    },
    enableNodes(node) {
      this.enableNodeMaps[node.id] = 1;
      node.disabled = false;
      if (node.children) {
        for (let i = 0; i < node.children.length; i++) {
          this.enableNodes(node.children[i])
        }
      }
    },
    removeDisabledNode() {
      let tmp = []
      for (let i = 0; i < this.audiencesTreeSelection.length; i++) {
        if (!this.enableNodeMaps[this.audiencesTreeSelection[i].id]) {
          tmp.push(this.audiencesTreeSelection[i])
        }
      }
      this.audiencesTreeSelection = tmp
    },
    disableNodes(node) {
      this.disableNodeArray.push(node)
      node.disabled = true;
      if (node.children) {
        for (let i = 0; i < node.children.length; i++) {
          this.disableNodes(node.children[i])
        }
      }
    },
    showSelectMaterialModal(){
      this.insertMaterialDialog = true;
    },
    insertImageToContent(url){
      window.tinymce.get('tinymce-graphics-'+this.materialId).focus();
      window.tinymce.get('tinymce-graphics-'+this.materialId).selection.setContent('<img src="'+url+'" />')
      this.insertMaterialDialog = false;
    },
    insertTemplateToContent(content){
      window.tinymce.get('tinymce-graphics-'+this.materialId).focus();
      window.tinymce.get('tinymce-graphics-'+this.materialId).selection.setContent(content)
      this.insertMaterialDialog = false;
    },
    destroyEditor() {
      this.showEditor = false;
    },
    editorInit() {
      this.showEditor = true;
    },
    getChannelTemplates() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/list', {type: 'template', channel_id: this.channel, name: ""})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.channelTemplates = [];
              this.templates = resData.data || [];
              for (let i = 0; i < this.templates.length; i++) {
                this.channelTemplates.push({
                      title: resData.data[i].name,
                      description: resData.data[i].abstract,
                      content: resData.data[i].content,
                    }
                )
              }
              this.destroyEditor();
              this.$nextTick(() => this.editorInit());
            } else {
              this.$toast('获取频道模版出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取频道模版出错：服务器出错！', {
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
      this.$graphicsHttp('post', '/material/image/upload', formData)
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
    notBeforeToday(date) {
      return date < moment().startOf('day').toDate();
    },
    storeAndPublish() {
      this.store(true);
    },
    getChannelAudience() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/channel/detail', {id: this.channel})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.audiencesTree[0].children = JSON.parse(resData.data.audiences);
              this.channelName = resData.data.name;
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

    getChannel() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/channels')
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
            this.$toast('获取频道出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    getImages() {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/list', {type: 'image', channel_id: this.channel, name: ''})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.images = resData.data || [];
            } else {
              this.$toast('获取图片出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取图片出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);

    },
    getMaterial() {
      if (this.materialId == 0) {
        return;
      }
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/detail', {id: this.materialId})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.content = resData.data.content;
              this.materialName = resData.data.name;
              this.imageUrl = resData.data.url;
              this.abstract = resData.data.abstract;
              this.channel = resData.data.channel_id;
              this.channelDisabled = true;
            } else {
              this.$toast('获取素材出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取素材出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },
    store(showPublish = false) {
      this.overlay += 1;
      this.$graphicsHttp('post', '/material/store', {
        name: this.materialName,
        content: this.content,
        abstract: this.abstract,
        url: this.imageUrl,
        id: this.materialId,
        channel_id: this.channel,
        type: 'graphics',
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
          if (showPublish === true) {
            document.body.scrollTop = document.documentElement.scrollTop = 0;
            this.publishDialog = showPublish;
          }

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
          .finally(() => this.overlay -= 1);
    },
    publish() {
      this.publishDialog = false;
      this.overlay += 1;
      let receivers = [];
      this.audiencesSelection.forEach(function (element) {
        receivers.push({id: element.id, type: element.type, name: element.name})
      });

      this.$graphicsHttp('post', '/material/publish', {
        id: this.materialId,
        receivers: JSON.stringify(receivers),
        send_time: this.sendTime,
        author: this.author,
      }).then(response => {
        console.log(JSON.stringify(response.data));
        let resData = response.data;
        if (resData.code === 0) {
          this.$toast("发布成功！", {
            type: 'success',
            timeout: 2000
          });
          location.href = '/message-owner'
        } else if (resData.code === 40 || resData.code === 1000004001) {
          this.$toast('发布出错：' + resData.message + '【' +
              resData.data.errors[Object.keys(resData.data.errors)[0]][0] + '】', {
            type: 'error',
          });
        } else {
          this.$toast('发布出错：' + resData.message, {
            type: 'error',
          });
        }
      })
          .catch(error => {
            console.log(error);
            this.$toast('发布出错：服务器出错！', {
              type: 'error',
              timeout: 2000,
            });
          })
          .finally(() => this.overlay -= 1);
    },

    handleAudiencesSearch: function (val) {
      if (val) {
        if (!this.audiencesAllOpened) {
          this.audiencesLastOpen = this.audiencesOpen;
          this.audiencesAllOpened = true;
          this.$refs.audiencesTree.updateAll(true);
        }
      } else {
        this.$refs.audiencesTree.updateAll(false);
        this.audiencesAllOpened = false;
        this.audiencesOpen = this.audiencesLastOpen;
      }
    }
  },
}
</script>
<style>
.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}

.tox-tinymce {
  border-bottom: none !important;
  border-left: none !important;
  border-right: none !important;
}

.my-date-picker-svg svg {
  margin-bottom: 0rem !important
}

</style>
