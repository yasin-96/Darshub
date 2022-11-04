<template>
  <footer class="footer p-10 bg-base-200">
    <div class="">
      <div class="avatar">
        <div class="w-24 rounded-full">
          <img width="128" src="logoUrl" alt="Logo" />
        </div>
      </div>
      <p>Copyright &copy; {{ new Date().getFullYear() }} - {{ $t("footer.cr") }}</p>
    </div>
    <div></div>
    <div>
      <!-- <span class="text-bold text-lg">Legal</span> 
    <a class="link link-hover">Terms of use</a> 
    <a class="link link-hover">Privacy policy</a> 
    <a class="link link-hover">Cookie policy</a> -->
    </div>
  </footer>
</template>

<script setup>
import { useI18n } from "vue-i18n";
const { t } = useI18n();

const logoFileName = "";
</script>

<!-- <style>
.footer {
  color: white
}
</style> -->
