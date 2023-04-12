<script lang="ts" setup>
export interface CourseInfo {
  id: String;
  name: String;
  description: String;
  duration: Date | String | null;
  level: String;
  author: String;
  released: Date | String | null;
  lastupdate: Date | String | null;
}

const props = defineProps<CourseInfo>();
</script>

<template>
  <v-card class="mx-auto" max-width="344">
    <v-img
      src="https://cdn.vuetifyjs.com/images/cards/sunshine.jpg"
      height="200px"
      cover
    ></v-img>

    <v-card-title>
      {{ props.name }}
    </v-card-title>

    <v-card-text>
      <div>{{ props.description }}</div>

      <div>Whitsunday Island, Whitsunday Islands</div>
    </v-card-text>

    <v-card-subtitle> {{ props.author }}</v-card-subtitle>

    <v-card-actions>
      <RouterLink :to="`/course/${props.id}/preview`">
        <v-btn color="orange"> Explore </v-btn>
      </RouterLink>
    </v-card-actions>
  </v-card>
</template>
