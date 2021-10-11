<template>
  <v-container>
    <v-dialog
        v-model="dialog"
        max-width="800px"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
            color="primary"
            dark
            class="mb-2"
            v-bind="attrs"
            v-on="on"
        >
          <v-icon>mdi-plus-thick</v-icon>菜单
        </v-btn>
      </template>
      <v-card>
        <v-card-title>
          <span class="headline">新增菜单</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                  cols="12"
                  sm="6"
                  md="3"
              >
                <v-text-field
                    v-model="editedItem.parent"
                    label="父级菜单id"
                ></v-text-field>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
                  md="3"
              >
                <v-text-field
                    v-model="editedItem.name"
                    label="名称"
                ></v-text-field>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
                  md="3"
              >
                <v-text-field
                    v-model="editedItem.icon"
                    label="图标"
                ></v-text-field>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
                  md="3"
              >
                <v-text-field
                    v-model="editedItem.route"
                    label="路由"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
              color="blue darken-1"
              text
              @click="close"
          >
            关闭
          </v-btn>
          <v-btn
              color="blue darken-1"
              text
              @click="save"
          >
            保存
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-treeview
        v-model="tree"
        :open="initiallyOpen"
        :items="items"
        activatable
        item-key="name"
        open-on-click
    >
      <template v-slot:prepend="{ item}">
        <v-icon v-if="!item.icon">
         mdi-folder
        </v-icon>
        <v-icon v-else>
          {{ item.icon }}
        </v-icon>
      </template>
      <template v-slot:append="{ item}">
        <v-btn @click="edit(item.id)" >
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
      </template>
    </v-treeview>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    editedItem: {
      parent: "",
      name: "",
      icon: "",
      route:""
    },
    defaultItem: {
      parent: "",
      name: "",
      icon: "",
      route:""
    },
    dialog: false,
    initiallyOpen: ['public'],
    tree: [],
    items:[],
  }),
  methods: {
    edit(id) {
      console.log(id)
    },
    save(){
      this.overlay = true;
      this.$graphicsHttp('post', '/menu/edit', this.editedItem).then((response) => {
        let resData = response.data;
        if (resData.code === 0) {
          this.channelTemplates = resData.data;
          this.destroyEditor();
          this.$nextTick(() => this.editorInit());
        } else {
          this.$toast('保存菜单出错：' + resData.message, {
            type: 'error',
          });
        }
      }).catch(error => {
        console.log(error);
        this.$toast('保存菜单出错：服务器出错！', {
          type: 'error',
          timeout: 2000,
        });
      }).finally(() => {
        this.overlay -= 1;
      });
      this.close()
    },
    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },
  },
  watch: {
    dialog(val) {
      val || this.close()
    },
  },
  created() {
    this.$graphicsHttp('post', '/menu/list').then((response) => {
      this.items = response.data;
    }).catch(error => {
      console.log(error);
    });
  }
}
</script>