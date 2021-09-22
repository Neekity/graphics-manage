<template>
  <editor :id="tinymceId" :init="init">
  </editor>
</template>

<script>
import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/icons/default/icons.js'
import 'tinymce/plugins/image'
import 'tinymce/plugins/link'
import 'tinymce/plugins/table'
import 'tinymce/plugins/advlist'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/hr'
import 'tinymce/plugins/template'
import 'tinymce/plugins/autoresize'
import 'tinymce/plugins/fullscreen'
import 'tinymce/plugins/paste'
import 'tinymce/plugins/emoticons'
import 'tinymce/plugins/autosave'

export default {
  name: 'TinymceMessageEditor',
  props: {
    id: {
      type: String,
      default: function () {
        return 'vue-tinymce-' + +new Date() + ((Math.random() * 1000).toFixed(0) + '')
      }
    },
    value: {
      type: String,
      default: ''
    },
    baseUrl: {
      type: String,
      default: ''
    },
    channel: [Number, String],
    editorTemplates: {
      type: Array,
      default: function (){
        return [];
      },
    },
  },
  watch: {
    value(val) {
      if (!this.hasChange && this.hasInit) {
        this.$nextTick(() =>
            window.tinymce.get(this.tinymceId).setContent(val || ''))
      }
    }
  },
  data() {
    const _this = this
    return {
      tinymceId: this.id,
      hasInit: false,
      hasChange: false,
      init: {
        base_url: this.baseUrl,
        toolbar_mode:'wrap',
        autosave_ask_before_unload:false,
        autosave_prefix: 'tinymce-autosave-{path}{query}-{id}-',
        toolbar_location:'auto',
        contextmenu_avoid_overlap: '.mce-spelling-word',
        contextmenu: "bold link image imagetools table spellchecker",
        skin_url: this.baseUrl + 'skins/ui/oxide', // skin路径
        language: 'zh_CN',
        language_url: this.baseUrl + 'langs/zh_CN.js',
        width:'100%',
        toolbar: 'restoredraft | forecolor backcolor bold italic link hr| bullist numlist table image template fullscreen emoticons | alignleft aligncenter alignright alignjustify outdent indent | formatselect fontselect fontsizeselect',
        content_style: "img {display: block;max-width: 100%; height: auto;} body{margin:0;}",
        image_description: false,
        placeholder: '请在此输入正文内容',
        link_default_protocol: "https",
        default_link_target: "_blank",
        toolbar_sticky: true,
        branding: false, // 去水印
        elementpath: false, // 禁用编辑器底部的状态栏
        statusbar: false, // 隐藏编辑器底部的状态栏
        paste_data_images: false, // 不允许粘贴图像
        menubar: false, // 隐藏最上方menu
        plugins: 'image link hr template autoresize table advlist lists fullscreen paste autosave emoticons', // 图片插件
        templates: this.editorTemplates,
        fontsize_formats: '12px 14px 16px 18px 24px 36px 48px 56px 72px',
        font_formats: '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;Comic Sans MS=comic sans ms,sans-serif;Courier New=courier new,courier;Georgia=georgia,palatino;Helvetica=helvetica;Impact=impact,chicago;Symbol=symbol;Tahoma=tahoma,arial,helvetica,sans-serif;Terminal=terminal,monaco;Times New Roman=times new roman,times;Verdana=verdana,geneva;Webdings=webdings;Wingdings=wingdings,zapf dingbats;知乎配置=BlinkMacSystemFont, Helvetica Neue, PingFang SC, Microsoft YaHei, Source Han Sans SC, Noto Sans CJK SC, WenQuanYi Micro Hei, sans-serif;小米配置=Helvetica Neue,Helvetica,Arial,Microsoft Yahei,Hiragino Sans GB,Heiti SC,WenQuanYi Micro Hei,sans-serif',
        images_upload_handler: (blobInfo, success, fail) => {
          let file = blobInfo.blob();//转化为易于理解的file对象
          let formData = new FormData();
          formData.append('graphics_img', file, file.name);
          formData.append('channel_id', _this.channel);
          _this.$http.post('/graphics/material/owner/image/upload', formData)
              .then(response => {
                console.log(JSON.stringify(response.data));
                let resData = response.data;
                if (resData.code === 0) {
                  success(resData.data.url)
                } else {
                  fail('上传失败：' + resData.message)
                }
              })
              .catch(error => {
                console.log(error);
                _this.$toast('上传出错：服务器出错！', {
                  type: 'error',
                  timeout: 2000,
                });
              });
        }, // 上传本地图片
        init_instance_callback: editor => {
          if (_this.value) {
            editor.setContent(_this.value)
          }
          _this.hasInit = true
          // 监听富文本内容的改变 将变化传回richtext的v-model
          editor.on('NodeChange Change SetContent', () => {
            if (!editor.getContent()){
              return;
            }
            this.hasChange = true
            this.$emit('input', editor.getContent())
          });
          editor.on('KeyUp', () => {
            this.hasChange = true
            this.$emit('input', editor.getContent())
          })
        },
      }
    }
  },
  components: {Editor},
  mounted() {
    tinymce.init({})
  },
}
</script>

<style scoped></style>
