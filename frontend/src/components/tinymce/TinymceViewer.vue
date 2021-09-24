<template>
  <editor disabled="disabled" :id="tinymceId" :init="init">
  </editor>
</template>

<script>
import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/icons/default/icons.js'
import 'tinymce/plugins/autoresize'

export default {
  name: 'TinymceViewer',
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
  },
  data() {
    const _this = this;
    return {
      tinymceId: this.id,
      hasInit: false,
      hasChange: false,
      init: {
        base_url: this.baseUrl,
        skin_url: this.baseUrl + 'skins/ui/oxide', // skin路径
        width: '100%',
        readonly: true,
        content_style: "img {display: block;max-width: 100%; height: auto;} body{margin:0;}",
        plugins: 'autoresize',
        branding: false, // 去水印
        menubar: false, // 隐藏最上方menu
        statusbar: false, // 隐藏编辑器底部的状态栏
        toolbar: [],
        init_instance_callback: editor => {
          if (_this.value) {
            editor.setContent(_this.value)
          }
        }
      }
    }
  },
  components: {Editor},
  created() {
    tinymce.init({})
  }
}
</script>
