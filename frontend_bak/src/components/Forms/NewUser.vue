<template>
  <div>
    <v-card>
      <v-card-text>
        <v-text-field outlined v-model="newUser.fullName" label="Full Name" placeholder="John Doe" required></v-text-field>
        <v-text-field outlined v-model="newUser.email" label="E-Mail" placeholder="John Doe" required></v-text-field>
        <v-text-field outlined v-model="newUser.nickName" label="Nick Name" placeholder="John Doe" required></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-btn text> Cancel </v-btn>
        <v-spacer></v-spacer>
        <v-btn color="primary" text @click="submit"> Submit </v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
export default {
  name: "NewUser",
  data: () => ({
    newUser: {
      fullName: null,
      email: null,
      nickName: null,
    },
  }),
};
</script>

<style></style>
