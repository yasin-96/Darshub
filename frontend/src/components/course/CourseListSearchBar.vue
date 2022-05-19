<template>
  <div  >

      <v-text-field outlined solo placeholder="Suche ..." :append-icon="isFilterEnabled?'mdi-filter-outline' : 'mdi-filter-off-outline'">

      </v-text-field>
  </div>
</template>

<script>
export default {
    name: "CourseListSearchBar",
    data: ()=>({
        isFilterEnabled: false
    })
}
</script>

<style>

</style>