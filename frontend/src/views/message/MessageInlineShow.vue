<template>
  <v-app>
    <v-card max-width="578px" class="mx-auto">
      <v-card-title>
        {{ title }}
      </v-card-title>
      <v-card-subtitle>
        {{subtitle}}
      </v-card-subtitle>

      <v-card-text class="mt-6">
        <v-row
            align="center"
            justify="center"
            class="mx-0 tinymce-content"
        >
          <tiny-mce-inline-viewer v-model="content" :skin-url="skinUrl"></tiny-mce-inline-viewer>
        </v-row>
      </v-card-text>

    </v-card>
  </v-app>
</template>

<script>
import TinyMceInlineViewer from "../../components/tinymce/TinyMceInlineViewer";

export default {
  name: "MessageInlineShow",
  components: {TinyMceInlineViewer},
  data() {
    return {
      skinUrl:'/tinymce/skins/ui/oxide',
      title: '我是标题',
      subtitle: '我是作者 2011-11-11',
      content: '<p><img src="https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/website_config/2020-05-19/19578.jpeg?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&amp;Expires=1663311317&amp;Signature=ttJWli0STvpNsEjWjNFdm3Gx%2F38%3D" alt="" width="2800" height="1750" /></p>\n' +
          '<p style="letter-spacing: 0.1px;"><img src="https://images.pexels.com/photos/1092730/pexels-photo-1092730.jpeg?auto=compress&amp;cs=tinysrgb&amp;dpr=1&amp;w=500" alt="" width="689" height="517" /></p>\n' +
          '<p style="letter-spacing: 0.1px;"><span style="font-weight: bolder;"><span style="font-size: 36px;">请在此输入正文内</span></span>容</p>\n' +
          '<table style="border-collapse: collapse; width: 300.531px;" border="10">\n' +
          '<tbody>\n' +
          '<tr>\n' +
          '<td style="width: 49.9219px;">1</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">4</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">44</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '</tr>\n' +
          '<tr>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">2</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">6</td>\n' +
          '</tr>\n' +
          '<tr>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">3</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">5</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '</tr>\n' +
          '<tr>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">45</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '<td style="width: 49.9219px;">&nbsp;</td>\n' +
          '</tr>\n' +
          '</tbody>\n' +
          '</table>\n' +
          '<p><span style="font-family: \'Microsoft YaHei\', \'Helvetica Neue\', \'PingFang SC\', sans-serif;">请</span></p>\n' +
          '<p><span style="font-family: \'Microsoft YaHei\', \'Helvetica Neue\', \'PingFang SC\', sans-serif;">asdf</span></p>\n' +
          '<p><span style="font-family: \'Microsoft YaHei\', \'Helvetica Neue\', \'PingFang SC\', sans-serif;">aa在</span></p>\n' +
          '<p><span style="font-family: \'Microsoft YaHei\', \'Helvetica Neue\', \'PingFang SC\', sans-serif;">12123</span></p>\n' +
          '<p><span style="font-family: \'Microsoft YaHei\', \'Helvetica Neue\', \'PingFang SC\', sans-serif;">此输入正文内容</span></p>\n' +
          '<p>&nbsp;</p>',
    }
  }
}
</script>

<style>
.tinymce-content img {
  display: block;
  width: 100%;
  height: auto;
}
</style>