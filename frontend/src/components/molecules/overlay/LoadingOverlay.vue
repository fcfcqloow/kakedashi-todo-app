<script lang="ts" setup>
type Props = {
  modelValue: boolean;
};

const props = defineProps<Props>();
</script>

<template>
  <v-overlay :model-value="props.modelValue" class="align-center justify-center"><v-progress-circular color="primary" indeterminate size="64" /></v-overlay>
</template>
