<template>
  <div>
    <v-card elevation="1">
      <v-data-table :items="loadedCourses" :headers="headers">
        <template v-slot:[`body.prepend`] class="no-hover d-sm-none d-md-flex">
          <tr class="pa-0">
            <td :colspan="headers.length" class="pa-1 ma-0">
              <v-text-field
                placeholder="Suche"
                outlined
                dense
                class="text-small ma-0"
              ></v-text-field>
            </td>
          </tr>
        </template>
        <template v-slot:[`item.ID`]="{ item }">
          <v-btn icon text small @click="editCourse(item.ID)">
            <v-icon color="success"> mdi-pencil </v-icon>
          </v-btn>
          {{ item.id }}
        </template>
      </v-data-table>
    </v-card>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  name: "AdminCourseList",
  computed: {
    ...mapState({
      loadedCourses: (state) => state.courseStore.general.allCourses,
    }),
    headers() {
      return [
        {
          text: "ID",
          align: "start",
          sortable: false,
          value: "ID",
          width: "80",
        },
        {
          text: "Name",
          align: "start",
          sortable: false,
          value: "Name",
        },
        {
          text: "Description",
          align: "start",
          sortable: false,
          value: "Description",
        },
        {
          text: "Duration",
          align: "start",
          sortable: false,
          value: "Duration",
        },
        {
          text: "Level",
          align: "start",
          sortable: false,
          value: "Level",
        },
        {
          text: "Author",
          align: "start",
          sortable: false,
          value: "Author",
        },
        {
          text: "Released",
          align: "start",
          sortable: false,
          value: "Released",
        },
        {
          text: "LastUpdate",
          align: "start",
          sortable: false,
          value: "LastUpdate",
        },
      ];
    },
  },
  methods: {
    editCourse(courseID) {
      this.$store.dispatch("courseStore/edit/act_setCourseIdForEditing", courseID)
    },
  },
};
</script>

<style></style>
