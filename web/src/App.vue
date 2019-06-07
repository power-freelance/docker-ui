<template>
  <v-app>
    <Navbar />

    <v-content>
      <Progress />
      <router-view/>
    </v-content>
  </v-app>
</template>

<script>
import Navbar from "@/components/organisms/Navbar";
import Progress from "@/components/atoms/Progress"

export default {
  components: {
    Navbar,
    Progress
  },
  data () {
    return {
      //
    }
  }
}
</script>
