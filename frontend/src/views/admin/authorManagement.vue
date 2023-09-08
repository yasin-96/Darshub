<script lang="ts" setup>
import { useAuthorCourseStore } from "@/stores/course/authorCourseStore";
import { computed, ref} from "vue";

const coreCourseStore = useAuthorCourseStore();

async function loadCourses() {
  await coreCourseStore.act_loadAllCourseAsQuickInfo();
}

const tHeaders = ref([
  {
    title: "sites.admin.courseManagement.table.header.id",
    align: "start",
    sortable: false,
    key: "id",
    value: "id",
  },
  {
    title: "sites.admin.courseManagement.table.header.name",
    align: "start",
    sortable: false,
    key: "name",
    value: "name",
  },
  {
    name: "sites.admin.courseManagement.table.header.description",
    align: "start",
    sortable: false,
    key: "description",
    value: "description",
  },
  {
    name: "sites.admin.courseManagement.table.header.author",
    align: "start",
    sortable: true,
    key: "author",
    value: "author",
  },
  {
    name: "sites.admin.courseManagement.table.header.duration",
    align: "start",
    sortable: false,
    key: "duration",
    value: "duration",
  },
  {
    name: "sites.admin.courseManagement.table.header.level",
    align: "start",
    sortable: false,
    key: "level",
    value: "level",
  },
  {
    name: "sites.admin.courseManagement.table.header.released",
    align: "start",
    sortable: false,
    key: "released",
    value: "released",
  },
  {
    name: "sites.admin.courseManagement.table.header.lastupdate",
    align: "start",
    sortable: false,
    key: "laskeytupdate",
    value: "laskeytupdate",
  },
  {
    name: "sites.admin.courseManagement.table.header.action",
    align: "start",
    sortable: false,
  },
]);
</script>

<template>
  <v-container>
    <v-card elevation="3">
      <v-data-table
        :items-per-page="5"
        :headers="tHeaders"
        :items-length="coreCourseStore.allCourseInMinimalForm.length"
        :items="coreCourseStore.allCourseInMinimalForm"
        class="elevation-1"
        @update:options="loadCourses"
      ></v-data-table>
    </v-card>
    <pre>{{coreCourseStore.allCourseInMinimalForm}}</pre>
  </v-container>
</template>
