<template>
  <div>
    <v-container fluid>
      <v-row>
        <v-col cols="4" class="mx-auto">
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "View_CourseDetails",
  components: {},
  computed: {
    ...mapState({
      loadedCourses: (state) => state.courseStore.general.allCourses,
    }),
  },
};
</script>

<style></style>
