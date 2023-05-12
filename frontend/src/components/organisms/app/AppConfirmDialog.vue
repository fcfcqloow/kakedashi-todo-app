<script lang="ts" setup>
import { useAppStore } from '@/stores/app';
import ConfirmDialog from '@/components/molecules/dialog/ConfirmDialog.vue';
const { confirm, closeConfirmDialog } = useAppStore();
</script>


<template>
  <div class="text-center">
    <confirm-dialog
      :model-value="confirm.isShow"
      :title="confirm.title"
      @update:model-value="!$event && closeConfirmDialog();"
      @ok="confirm.ok(); closeConfirmDialog();"
      @cancel="confirm.cancel(); closeConfirmDialog();"
    />
  </div>
</template>