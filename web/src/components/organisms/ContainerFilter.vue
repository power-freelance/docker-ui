<template>
  <v-form>
    <ContainerStatusInput multiple v-model="status" />
  </v-form>
</template>

<script>
import { mapFields } from 'vuex-map-fields'

import ContainerStatusInput from "@/components/atoms/ContainerStatusInput";

export default {
  components: {
    ContainerStatusInput
  },
  computed: {
    ...mapFields('containers/list', {
      status: 'filter.status'
    })
  }
}
</script>

<style scoped>

</style>