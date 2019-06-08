<template>
  <tr>
    <td>{{Id}}</td>
    <td>{{Names}}</td>
    <td>{{Image}}</td>
    <td>{{State}}</td>
    <td>{{Created}}</td>
  </tr>
</template>

<script>
export default {
  props: {
    Id: String,
    Names: Array,
    Image: String,
    State: String,
    Created: Number
  }
}
</script>
