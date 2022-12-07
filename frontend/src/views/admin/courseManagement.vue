<script lang="ts" setup>
import { useCoreCourseStore } from "@/stores/course/coreCourseStore";
import { computed } from "vue";

const coreCourseStore = useCoreCourseStore();

const tHeaders = computed(() => {
  return [
    {
      name: "sites.admin.courseManagement.table.header.id",
    },
    {
      name: "sites.admin.courseManagement.table.header.name",
    },
    {
      name: "sites.admin.courseManagement.table.header.description",
    },
    {
      name: "sites.admin.courseManagement.table.header.duration",
    },
    {
      name: "sites.admin.courseManagement.table.header.level",
    },
    {
      name: "sites.admin.courseManagement.table.header.content",
    },
    {
      name: "sites.admin.courseManagement.table.header.author",
    },
    {
      name: "sites.admin.courseManagement.table.header.released",
    },
    {
      name: "sites.admin.courseManagement.table.header.lastupdate",
    },
    {
      name: "sites.admin.courseManagement.table.header.action",
    },
  ];
});
</script>

<template>
  <div class="">
    <table class="shadow-lg rounded-lg">
      <thead class="border-b-2">
        <tr class="text-left uppercase">
          <th
            v-for="(th, index) in tHeaders"
            :key="`user-manager-h-${index}`"
            :class="`p-2 ${
              index == 0 || tHeaders.length - 1
                ? 'border-l-0 border-r-0'
                : 'border-l-4 border-r-4'
            }`"
          >
            {{ $t(th.name) }}
          </th>
        </tr>
      </thead>
      <tbody class="">
        <tr
          v-for="(td, index) in coreCourseStore.allCourseInMinimalForm"
          :key="`user-manager-d-${index}`"
        >
          <td class="p-2">{{ td.id }}</td>
          <td class="p-2">{{ td.name }}</td>
          <td class="p-2">{{ td.description }}</td>
          <td class="p-2">{{ td.duration }}</td>
          <td class="p-2">{{ td.level }}</td>
          <td class="p-2">{{ td.content }}</td>
          <td class="p-2">{{ td.author }}</td>
          <td class="p-2">{{ td.released }}</td>
          <td class="p-2">{{ td.lastupdate }}</td>
          <td class="p-2">
            <button
              class="rounded-full hover:shadow-xl hover:bg-slate-200 hover:bg-slate-400 px-2 py-2"
            >
              <i class="i-heroicons-pencil text-red-600 block"></i>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
