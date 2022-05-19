<template>
  <div>
      <v-container fluid>
      <v-row>
       <v-col>
           <clsearchbar/>
       </v-col>
      </v-row>
    </v-container>
    <v-container fluid>
      <v-row>
        <v-col cols="12" sm="12" md="12" lg="8" xl="6" class="mx-auto" 
        v-for="course in loadedCourses" :key="course.id"
        > 
          <mcd :cTitle="course.name" :cDescription="course.description" />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import MinimalCourseDetails from "@/components/course/MinimalCourseDetails"
import CourseListSearchBar from "@/components/course/CourseListSearchBar"
export default {
  name: "View_CourseOverview",
  components:{
    mcd: MinimalCourseDetails,
    clsearchbar: CourseListSearchBar
  },
  computed:{
      ...mapState({
          loadedCourses: (state) => state.ss
      })

  }
};
</script>

<style></style>
