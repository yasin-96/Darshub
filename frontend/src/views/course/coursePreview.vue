<script lang="ts" setup>
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { computed } from "vue";

const coreCourseStore = useCoreCourseStore();

const coursePreview = computed(() => {
  return coreCourseStore.getCoursePreview;
});
</script>

<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>
          {{ coursePreview.name }}
        </h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <h2>Description</h2>
      </v-col>
      <v-col> {{ coursePreview.description }} </v-col>
    </v-row>
    <v-row class="pt-10">
      <v-col>
        <h3 class="">Reviews</h3>
        TODO
        <v-rating
          :v-model="coursePreview.level"
          color="yellow"
          bg-color="orange-lighten-1"
          readonly
        ></v-rating>
      </v-col>
    </v-row>
    <v-row class="pt-10">
      <v-col>
        <p>TODO Alle Kapitel auflisten</p>

        <pre>{{ coursePreview }}</pre>
      </v-col>
    </v-row>
  </v-container>
</template>
