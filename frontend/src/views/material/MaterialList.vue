<template>
  <v-app>
    <v-card>
      <v-card-text>
        <v-tabs v-model="tab">
          <v-tabs-slider></v-tabs-slider>
          <v-tab
              v-for="item in items"
              :key="item"
          >{{ item }}
          </v-tab>
        </v-tabs>
        <v-tabs-items v-model="tab">
          <v-tab-item
              key="图文素材">
            <material-graphics-list></material-graphics-list>
          </v-tab-item>
          <v-tab-item
              key="图片">
            <image-list></image-list>
          </v-tab-item>
          <v-tab-item
              key="模版">
            <material-template-list ></material-template-list>
          </v-tab-item>
        </v-tabs-items>
      </v-card-text>
    </v-card>
  </v-app>
</template>

<script>
import ImageList from "../../components/material/image/ImageList";
import MaterialGraphicsList from "../../components/material/graphics/MaterialGraphicsList";
import MaterialTemplateList from "../../components/material/template/MaterialTemplateList";

export default {
  name: "MaterialList",
  components: {MaterialGraphicsList, ImageList,MaterialTemplateList},
  props:['skinUrl'],
  data() {
    return {
      items: ['图文', '图片', '模版'],
      tab: null,
    }
  },
}
</script>
