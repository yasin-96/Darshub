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
  <div class="card card-compact bg-base-100 shadow-xl">
    <figure>
      <img
        src="https://images.pexels.com/photos/68562/pexels-photo-68562.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
        alt="Placeholder"
        class="h-full w-full object-cover object-center group-hover:opacity-75"
      />
    </figure>
    <div class="card-body">
      <h2 class="card-title">{{ props.name }}</h2>
      <p>
        {{ props.description }}
      </p>
      <p>
        {{ props.level }}
      </p>
      <div class="card-actions justify-end">
        <a :href="`/course/${props.id}/preview`" class="group link">
          <button class="btn btn-primary">Open</button>
        </a>
      </div>
    </div>
  </div>
</template>
