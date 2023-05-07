<script setup lang="ts">
import LoadingOverlay from '@/components/molecules/overlay/LoadingOverlay.vue';
import { useAppStore } from '@/stores/app';

const { isGlobalLoading } = useAppStore();
</script>

<template>
  <loading-overlay :model-value="isGlobalLoading" />
</template>
