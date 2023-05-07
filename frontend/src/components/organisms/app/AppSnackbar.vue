<script lang="ts" setup>
import { useAppStore } from '@/stores/app';
import Snackbar from '@/components/molecules/snackbar/Snackbar.vue';

const { snackbar, closeSnackbar } = useAppStore();

</script>
<template>
  <snackbar
    :model-value="snackbar.isShow"
    :timeout="4000"
    :color="snackbar.color"
    @update:model-value="closeSnackbar"
  >
    {{ snackbar.value }}
  </snackbar>
</template>
