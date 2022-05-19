<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card height="200">
          <v-card-title> {{ cTitle }} </v-card-title>
          <v-card-subtitle> </v-card-subtitle>
          <v-card-text>
            {{ cDescription }}
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "MinimalCourseDetails",
  props: {
    cTitle: {
      type: String,
      required: true,
    },
    cDescription: {
      type: String,
      required: true,
    },
  },
};
</script>

<style></style>
