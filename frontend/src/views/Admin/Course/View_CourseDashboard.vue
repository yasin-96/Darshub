<template>
  <div>
      <v-container fluid>
      <v-row>
       <v-col cols="12" class="mx-auto">
           <AdminCourseList />
       </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import AdminCourseList from "@/components/course/Admin/AdminCourseList"
export default {
  name: "View_CourseOverview",
  components:{
   AdminCourseList
  }
};
</script>

<style></style>
