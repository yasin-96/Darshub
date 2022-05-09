<template>
  <div class="about">
    <UserDetails />
    <UserCourseList />
  </div>
</template>

<script>
import UserDetails from "@/components/User/UserDetails";
import UserCourseList from "@/components/User/UserCourseList";
export default {
  name: "UserDashboard",
  components: {
    UserDetails,
    UserCourseList,
  },
};
</script>

<style></style>
