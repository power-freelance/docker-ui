<template>
  <v-data-table :items="items" :headers="headers">
    <template v-slot:items="{item}">
      <ContainerTableRow v-bind="item" />
    </template>
  </v-data-table>
</template>

<script>
import { mapState } from 'vuex'

import ContainerTableRow from "@/components/molecules/ContainerTableRow";

export default {
  components: {ContainerTableRow},
  data() {
    return {
      headers: [
        {text: 'ID', align: 'left', sortable: false, value: 'ID'},
        {text: 'Name', align: 'left', sortable: false, value: 'Name'},
        {text: 'Image', align: 'left', sortable: false, value: 'Image'},
        {text: 'Status', align: 'left', sortable: false, value: 'State'},
        {text: 'Created', align: 'left', sortable: false, value: 'Created'}
      ]
    }
  },
  computed: {
    ...mapState('containers/list', {
      items: state => state.items
    })
  }
}
</script>

<style scoped>

</style>