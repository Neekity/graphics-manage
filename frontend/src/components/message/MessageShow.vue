<template>
  <v-app>
    <v-card max-width="660px" width="100%" class="mx-auto" height="100%">
      <v-card-title>
        {{ title }}
      </v-card-title>
      <v-card-subtitle class="text-left">
        {{subtitle}}
      </v-card-subtitle>
      <v-card-text class="mt-4">
        <v-row class="mx-0">
          <tinymce-viewer v-model="content" :base-url="baseUrl"></tinymce-viewer>
        </v-row>
      </v-card-text>
    </v-card>
  </v-app>
</template>

<script>
import TinymceViewer from "../tinymce/TinymceViewer";
export default {
  name: "MessageShow",
  components: {TinymceViewer},
  props:['baseUrl','content','title','subtitle']
}
</script>


<style>
.tox-tinymce {
  border:none !important
}
</style>
