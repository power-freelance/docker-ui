<template>
  <VSelect
    solo
    hide-details
    multiple
    chips
    deletable-chips
    clearable
    :value="value"
    :items="items"
    item-value="id"
    item-text="text"
    @input="$emit('input', $event)"
  />
</template>

<script>
import items from '@/assets/data/container-status'

export default {
  props: {
    value: [Array, String],
    multiple: {type: Boolean, default: false}
  },
  data() {
    return {
      items: items
    }
  }
}
</script>

<style scoped>

</style>