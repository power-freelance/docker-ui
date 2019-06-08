<template>
  <v-layout class="containers-page" column>
    <v-flex shrink>
      <ContainerFilter />
    </v-flex>
    <v-flex shrink>
      <ContainerTable />
    </v-flex>
  </v-layout>
</template>

<script>
import { mapState, mapActions } from 'vuex'

import ContainerTable from "@/components/organisms/ContainerTable";
import ContainerFilter from "@/components/organisms/ContainerFilter";

export default {
  components: {
    ContainerFilter,
    ContainerTable
  },
  computed: {
    ...mapState('containers/list', {
      filter: state => state.filter
    })
  },
  watch: {
    filter: {
      deep: true,
      handler() {
        this.fetchItems();
      }
    }
  },
  mounted() {
    this.fetchItems();
  },
  methods: {
    ...mapActions('containers/list', ['fetchItems']),
  }
}
</script>
