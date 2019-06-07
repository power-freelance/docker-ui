<template>
  <v-layout row wrap class="dashboard-page">
    <v-flex xs6>
      <v-card v-if="info">
        <v-toolbar dark color="primary">
          <v-toolbar-title>
            Statistics
          </v-toolbar-title>
        </v-toolbar>
        <v-list two-line>
          <template v-for="(item, index) in items">
            <v-list-tile :key="index" avatar ripple :to="item.to" :exact="item.exact">
              <v-list-tile-content>
                <v-list-tile-title>{{item.name}}</v-list-tile-title>
                <v-list-tile-sub-title>{{item.description}}</v-list-tile-sub-title>
              </v-list-tile-content>
              <v-list-tile-action>
                <v-chip color="primary" class="font-weight-bold" label outline>
                  {{info.docker.info[item.key]}}
                </v-chip>
              </v-list-tile-action>
            </v-list-tile>
            <v-divider v-if="index + 1 < items.length" :key="`divider-${index}`"></v-divider>
          </template>
        </v-list>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'

export default {
  data() {
    return {
      items: [
        { name: 'Containers', description: 'Container count', key: 'Containers', to: '/'},
        { name: 'Images', description: 'Image count', key: 'Images', to: '/'}
      ]
    }
  },
  computed: {
    ...mapState('dashboard', {
      info: state => state.info
    })
  },
  mounted() {
    this.$store.dispatch('dashboard/fetchInfo');
  }
}
</script>

<style scoped>

</style>