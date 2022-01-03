<template>
  <v-app class="pa-9">
    <div class="pb-8"></div>
    <v-carousel v-model="model">
      <v-carousel-item :key="0">
        <v-sheet
            height="100%"
            tile
        >
          <v-row class="fill-height" align="center" justify="center">
            <div class="text-h4">
              一款开箱即用的图文消息管理web应用
            </div>
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                    icon
                    large
                    color="primary"
                    href="/material"
                    dark
                    v-bind="attrs"
                    v-on="on"
                >
                  <v-icon>mdi-location-enter</v-icon>
                </v-btn>
              </template>
              <span>点击开始</span>
            </v-tooltip>
          </v-row>
        </v-sheet>
      </v-carousel-item>
      <v-carousel-item :key="1">
        <v-sheet
            height="100%"
            tile
        >
          <v-row class="fill-height" align="center" justify="center">
              <h3>支持casbin权限控制</h3>
          </v-row>
          <v-row class="fill-height" align="center" justify="center">
            <h3>前后端分离</h3>
          </v-row>
          <v-row class="fill-height" align="center" justify="center">
            <h3>Oauth登陆</h3>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel>
  </v-app>

</template>

<script>
export default {
  name: "ShowMe",
  data() {
    return {
      model: 0,
    }
  }
}
</script>

<style scoped>

</style>