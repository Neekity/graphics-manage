<template>
  <v-row>
    <v-col cols="12"
           lg="2"
           md="3"
           sm="4"
           xs="6">
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
  </v-row>
</template>

<script>
export default {
  name: "ChannelSelect"
}
</script>

<style scoped>

</style>