<template>
  <v-progress-linear
    class="ma-0"
    indeterminate
    background-opacity="0"
    color="primary"
    background-color="primary"
    height="4"
    :active="value"
  />
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState('ui', {
      value: state => state.progress
    })
  }
}
</script>