<template>
  <v-layout class="containers-page" row wrap>
    <v-flex pa-2 xs12 sm6 v-for="(item, index) in items" :key="index">
      <v-card>
        <v-card-title>{{item.Names[0]}}</v-card-title>
        <v-card-text>
          <v-chip color="primary" label outline>
            <v-avatar>
              <v-icon>mdi-image</v-icon>
            </v-avatar>
            {{item.Image}}
          </v-chip>

          <v-chip color="black" label outline>
            <v-avatar>
              <v-icon>mdi-console</v-icon>
            </v-avatar>
            {{item.Command}}
          </v-chip>
        </v-card-text>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState('containers', {
      items: state => state.items
    })
  },
  mounted() {
    this.$store.dispatch('containers/fetchItems');
  }
}
</script>
