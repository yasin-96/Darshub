<script setup lang="ts">
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { useLoginStore } from "@/stores/session/loginStore";

const loginStore = useLoginStore();
const coreCourseStore = useCoreCourseStore();
</script>

<template>
  <v-card elevation="0">
    <v-container fluid>
      <v-row>
        <v-col>
          <v-btn icon color="green">
            <v-icon large>mdi-plus-circle-outline</v-icon>
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<style scoped></style>
