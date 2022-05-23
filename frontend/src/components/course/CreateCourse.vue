<template>
  <v-form
    ref="form"
    v-model="baseInfoValidation"
    lazy-validation
  >
    <v-text-field
      v-model="courseName"
      :counter="10"
      :rules="nameRules"
      label="Kurs Bezeichnung"
      required
    ></v-text-field>

    <v-text-field
      v-model="email"
      :rules="emailRules"
      label="E-mail"
      required
    ></v-text-field>


    <!-- <v-btn
      :disabled="!valid"
      color="success"
      class="mr-4"
      @click="validate"
    >
      Validate
    </v-btn> -->

    <!-- <v-btn
      color="error"
      class="mr-4"
      @click="reset"
    >
      Reset Form
    </v-btn> -->

    <!-- <v-btn
      color="warning"
      @click="resetValidation"
    >
      Reset Validation
    </v-btn> -->
  </v-form>
</template>

<script>
export default {
  name: "CreateCourse",
  data: ()=>({
    baseInfoValidation: null
  })
}
</script>

<style>

</style>