<script lang="ts" setup>
import { useAppStore } from '@/stores/app';
import RestoreButton from '@/components/atoms/RestoreButton.vue';

type Props = {
};
type Emit = {
  (e: 'click'): void;
};

const { settings, showConfirmDialog } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>
<template>
   <v-tooltip
    v-if="settings.sortable"
    text="Delete and Save"
    location="bottom"
  >
    <template v-slot:activator="{ props }">
      <restore-button
        v-bind="props"
        size="small"
        :style="`
          margin: 10px;
          background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
        `"
        @click="showConfirmDialog({ title : 'Delete and Save for Done?', ok : () => emits('click') }); "
      />
    </template>
  </v-tooltip>
</template>