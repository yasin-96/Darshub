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
  <a :href="`/course/${props.id}/preview`" class="group">
    <div class="w-full overflow-hidden rounded-lg bg-gray-200">
      <img
        src="https://images.pexels.com/photos/68562/pexels-photo-68562.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        alt="Placeholder"
        class="h-full w-full object-cover object-center group-hover:opacity-75"
      />
    </div>
    <h3 class="mt-4 text-sm text-gray-700">{{ props.name }}</h3>
    <p class="mt-1 text-lg font-medium text-gray-900">{{ props.level }}</p>
  </a>
</template>
