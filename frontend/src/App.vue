<template>
  <v-app>
    <v-navigation-drawer
        v-model="drawer"
        app>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            图文消息管理系统
          </v-list-item-title>
          <v-list-item-subtitle>
            v1.0
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-list
          dense
          nav
      >
        <v-list-group
            v-for="item in items"
            :key="item.id"
            v-model="item.active"
            :prepend-icon="item.icon || 'mdi-folder'"
            no-action
        >
          <template v-slot:activator>
            <v-list-item-content>
              <v-list-item-title v-text="item.name"></v-list-item-title>
            </v-list-item-content>
          </template>
          <v-list-item
              v-for="subItem in item.children"
              :key="subItem.id"
              link
              :input-value="subItem.active"
              :href="subItem.path"
          >
            <v-list-item-icon>
              <v-icon>{{ subItem.icon }}</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title v-text="subItem.name"></v-list-item-title>
            </v-list-item-content>

          </v-list-item>
        </v-list-group>
      </v-list>
      <!--  -->
    </v-navigation-drawer>

    <v-app-bar app>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>{{ title }}</v-toolbar-title>
      <v-spacer></v-spacer>
      <div>
        <v-menu v-if="user.userId"
                bottom
                min-width="150px"
                rounded
        >
          <template v-slot:activator="{ on }">
            <v-btn
                icon
                v-on="on"
            >
              <v-avatar>
                <img
                    :src="user.avatarImgPath"
                >
              </v-avatar>
            </v-btn>
          </template>
          <v-card>
            <v-list-item-content class="justify-center">
              <div class="mx-auto text-center">
                <h4 class="my-3">{{ user.userName }}</h4>
                <v-divider class="my-3"></v-divider>
                <v-btn
                    depressed
                    rounded
                    text
                    @click="logout"
                >
                  退出
                </v-btn>
              </div>
            </v-list-item-content>
          </v-card>
        </v-menu>
      </div>
    </v-app-bar><!-- 根据应用组件来调整你的内容 -->
    <v-main><!-- 给应用提供合适的间距 -->
        <router-view></router-view>
    </v-main>

    <v-footer app>
      <!-- -->
    </v-footer>
  </v-app>
</template>

<script>
export default {
  name: 'App',
  data: () => ({
    title: "",
    drawer: null,
    activeParentMenu: null,
    items: [],
    right: null,
    user: {}
  }),
  methods: {
    logout() {
      this.$store.dispatch('handleLogout', {}).then(() => {
        this.$router.push({
          path: '/login'
        });
      });
    },
  },
  mounted() {
    this.items = this.$store.state.user.menu;
    this.user = this.$store.state.user;
  },
}
</script>
