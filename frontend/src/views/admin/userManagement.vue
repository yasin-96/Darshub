<script lang="ts" setup>
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { computed } from "vue";

const coreCourseStore = useCoreCourseStore();

const tHeaders = computed(() => {
  return [
    {
      name: "",
    },
  ];
});
</script>

<template>
  <table class="table table-compact w-full">
    <thead>
      <tr>
        <th v-for="(th, index) in tHeaders" :key="`user-manager-h-${index}`">
          {{ th }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="(td, index) in coreCourseStore.allCourseInMinimalForm"
        :key="`user-manager-d-${index}`"
      >
        <td>{{td.id}}</td>
        <td>{{td.name}}</td>
        <td>{{td.description}}</td>
        <td>{{td.duration}}</td>
        <td>{{td.level}}</td>
        <td>{{td.content}}</td>
        <td>{{td.author}}</td>
        <td>{{td.released}}</td>
        <td>{{td.lastupdate}}</td>
      </tr>
    </tbody>
  </table>
</template>
