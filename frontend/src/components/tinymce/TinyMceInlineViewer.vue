<template>
  <editor disabled="disabled" :id="tinymceId" :init="init">
  </editor>
</template>

<script>
import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'
import 'tinymce/themes/silver'
import 'tinymce/icons/default/icons.js'

export default {
  name: 'TinyMceInlineViewer',
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
    skinUrl:{
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
        skin_url: this.skinUrl, // skin路径
        height: 300, // 编辑器高度
        readonly: true,
        inline: true,
        branding: false, // 去水印
        menubar: false, // 隐藏最上方menu
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
