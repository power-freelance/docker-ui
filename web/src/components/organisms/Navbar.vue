<template>
  <v-toolbar app dark color="primary">
    <v-toolbar-title>
      <v-avatar size="120">
        <Logo  />
      </v-avatar>
    </v-toolbar-title>

    <v-spacer></v-spacer>

    <v-toolbar-items>
      <v-btn flat title="Dashboard" :to="{name: 'dashboard'}" :exact="true">
        <v-icon>mdi-monitor-dashboard</v-icon>
      </v-btn>
      <v-btn flat title="Containers" :to="{name: 'containers'}">
        <v-icon>mdi-package-variant-closed</v-icon>
      </v-btn>
      <v-btn flat title="Images" :to="{name: 'images'}">
        <v-icon>mdi-image</v-icon>
      </v-btn>
    </v-toolbar-items>

  </v-toolbar>
</template>

<script>
import Logo from "@/components/atoms/Logo"

export default {
  components: {
    Logo
  }
}
</script>
