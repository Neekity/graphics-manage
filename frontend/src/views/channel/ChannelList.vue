<template>
  <v-app>
    <v-card class="my-2">
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 sm="6"
          >
            <v-text-field
                v-model="channelName"
                label="请输入标题"
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getChannel"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-spacer></v-spacer>
          <v-icon
              class="mr-2"
              color="blue darken-2"
              @click="edit(0)"
          >
            mdi-plus
          </v-icon>
        </v-row>
      </v-card-title>
      <v-data-table
          :headers="headers"
          :items="channelItems"
          item-key="id"
          outlined
          :footer-props="{
              showFirstLastPage: true,
              firstIcon: 'mdi-arrow-collapse-left',
              lastIcon: 'mdi-arrow-collapse-right',
              prevIcon: 'mdi-minus',
              nextIcon: 'mdi-plus'
            }"
      >
        <template v-slot:item.actions="{ item }">
          <v-icon
              small
              class="mr-2"
              color="blue darken-2"
              @click="edit(item.id)"
          >
            mdi-pencil
          </v-icon>
        </template>
      </v-data-table>
    </v-card>

  </v-app>
</template>

<script>
export default {
  name: "ChannelList",
  data() {
    return {
      headers: [{
        text: '频道名称',
        align: 'start',
        sortable: false,
        value: 'name',
      },
        {text: 'Actions', value: 'actions', sortable: false},
      ],
      channelItems: [],
      channelName:'',
    }
  },
  created() {
    this.getChannel()
  },
  methods: {
    edit(id){
      location.href='channel/edit?id='+id;
    },
    getChannel() {
      this.overlay += 1;
      this.$graphicsHttp('post','/channel/list',{name:this.channelName})
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
  }
}
</script>

<style scoped>

</style>