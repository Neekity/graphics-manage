<template>
  <v-app>
    <v-card>
      <v-card-title>
        <v-row class="pl-2">
          <v-col
              cols="12"
              lg="3"
              md="4"
              sm="6"
              xs="12"
          >
            <v-text-field
                v-model="title"
                label="请输入标题"
                single-line
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getUserMessage"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-col
              cols="12"
              lg="3"
              md="4"
              sm="6"
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
        </v-row>

      </v-card-title>
      <v-card-text>
        <v-lazy
            v-model="isActive"
            :options="{
          threshold: .5
        }"
            min-height="200"
            transition="fade-transition"
        >
          <v-container fluid>
            <v-row>
              <v-col
                  v-for="message in messages"
                  :key="message.id"
                  cols="12"
                  lg="3"
                  md="4"
                  sm="6"
                  xs="12"
              >
                <v-card @click="view(message.id)">
                  <v-img
                      :src="message.cover_picture_url"
                      aspect-ratio="2.347"
                  >
                  </v-img>
                  <v-card-title class="text-h6">
                    {{ message.title }}
                  </v-card-title>
                  <v-card-subtitle class="text-h6">
                    {{ message.abstract }}
                  </v-card-subtitle>
                </v-card>
              </v-col>
            </v-row>
          </v-container>
        </v-lazy>
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
  name: 'MessageMasonry',
  data() {
    return {
      isActive: false,
      messages: [],
      overlay: false,
      title: "",
      channel: null,
      channelItems: [{id: '1', name: '宫傲'}, {id: 2, name: '够阿斯顿'}]
    }
  },
  created() {
    this.getUserMessage();
  },
  watch: {
    channel: function () {
        console.log(this.channel);
    }
  },
  methods: {
    getUserMessage() {
      this.messages = [
        {
          id: 1,
          title: '图文素材1',
          abstract: '123asfaf123',
          cover_picture_url: 'https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/graphics/2021-09-08/19612-5480668e3dfe10f449bd01a62c2a6986?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&Expires=2261817152&Signature=o3B1qfWYkAul8KCI%2B4E%2FbF4jtdQ%3D',
          content: '<p>请在此输入正文<img src="https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/graphics/2021-09-08/19612-5480668e3dfe10f449bd01a62c2a6986?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&amp;Expires=2261817152&amp;Signature=o3B1qfWYkAul8KCI%2B4E%2FbF4jtdQ%3D" alt="" width="768" height="475" /><img src="https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/graphics/2021-09-08/19612-5480668e3dfe10f449bd01a62c2a6986?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&amp;Expires=2261817152&amp;Signature=o3B1qfWYkAul8KCI%2B4E%2FbF4jtdQ%3D" alt="" /></p>\n' +
              '<p>&nbsp;</p>\n' +
              '<p><img src="https://oa-test-knzn02.oss-cn-hongkong.aliyuncs.com/graphics/2021-08-23/19597-b9a531f8fcc8e864e1dee7be3ec088e5?OSSAccessKeyId=LTAI4Fv4DUBzjMg5qjWyudpN&amp;Expires=2260418962&amp;Signature=QawtpgimFQAoSAx%2B5N05VQNsXnI%3D" alt="" /></p>',
          send_time: '2021-11-11 12:23:12',
        }
      ];
    },
    view(id) {
      location.href = '/graphics/message/user/' + id;
    },
  },
}
</script>
