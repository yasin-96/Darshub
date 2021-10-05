<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <NewUser />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import NewUser from "../components/Forms/NewUser.vue";

export default {
  name: "Home",
  components: {
    NewUser,
  },
  data: () => ({}),
};
</script>
