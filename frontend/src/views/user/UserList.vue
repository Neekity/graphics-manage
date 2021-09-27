<template>
  <v-app>
    <v-card width="960" class="mx-auto my-2">
      <v-card-title>
        <v-row>
          <v-col cols="12"
                 sm="6"
          >
            <v-text-field
                v-model="userName"
                label="请输入名称"
                outlined
                dense
            >
              <v-icon
                  slot="append"
                  color="blue"
                  @click="getUsers"
              >
                mdi-magnify
              </v-icon>
            </v-text-field>
          </v-col>
          <v-spacer></v-spacer>
        </v-row>
      </v-card-title>
      <v-data-table
          :headers="headers"
          :items="userItems"
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
  name: "UserList",
  data() {
    return {
      headers: [{
        text: '用户名称',
        align: 'start',
        value: 'name',
      },
        {text: '用户邮箱', value: 'email'},
        {text: '用户电话', value: 'mobile'},
        {text: '操作', value: 'actions', sortable: false},
      ],
      userItems: [],
      userName:'',
    }
  },
  created() {
    this.getUsers()
  },
  methods: {
    edit(id){
      location.href='channel/edit?id='+id;
    },
    getUsers() {
      this.overlay += 1;
      this.$axios.post('/backend/user',{name:this.userName})
          .then(response => {
            let resData = response.data;
            if (resData.code === 0) {
              this.userItems = resData.data;
            } else {
              this.$toast('获取用户出错：' + resData.message, {
                type: 'error',
              });
            }
          })
          .catch(error => {
            console.log(error);
            this.$toast('获取用户出错：服务器出错！', {
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