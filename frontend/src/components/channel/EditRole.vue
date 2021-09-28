<template>
  <div>
    <v-card class="mx-auto">
      <v-card-text>
        <v-row>
          <v-col cols="12" md="4" sm="6">
            <v-text-field
                v-model="roleName"
                dense
                outlined
                label="角色名称"
            ></v-text-field>
          </v-col>
          <v-col cols="12" md="4" sm="6">
            <v-text-field
                v-model="casbinRole"
                dense
                outlined
                label="casbin角色key"
                :disabled="!!roleId"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row class="text-left">
          <v-col md="4" sm="6">
            <v-text-field
                v-model="search"
                label="输入关键字搜索用户"
                clearable
                dense
                outlined
                @input="handleSearch"
                clear-icon="mdi-close-circle-outline"
            ></v-text-field>
            <v-treeview
                ref="items"
                v-model="selection"
                :items="items"
                selectable
                return-object
                dense
                hoverable
                :open.sync="open"
                :search="search"
            ></v-treeview>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col md="4" sm="6">
            <div class="text-h6 pb-2">分配角色至用户</div>
            <v-scroll-x-transition
                group
                hide-on-leave
            >
              <v-chip
                  v-for="(node, i) in selection"
                  :key="i"
                  color="primary"
                  outlined
                  small
                  close
                  class="ma-1"
                  @click:close="selection.splice(i,1)"
              >
                <v-icon
                    left
                    small
                >
                  mdi-account
                </v-icon>
                {{ node.name }}
              </v-chip>
            </v-scroll-x-transition>
          </v-col>
        </v-row>
      </v-card-text>
      <v-card-actions>
        <v-btn
            color="success"
            @click="store"
            class="text-center"
        >
          保存
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-overlay :value="overlay">
      <v-progress-circular
          indeterminate
          color="primary"
      ></v-progress-circular>
    </v-overlay>
  </div>
</template>

<script>
export default {
  name: "EditRole",
  props: {
    roleId: {
      type: Number,
      default: 0,
    }
  },
  data() {
    return {

    }
  },
  mounted() {
    this.getAllUsers();
  },
  methods: {

  }
}
</script>

<style scoped>

</style>