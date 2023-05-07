<script lang="ts" setup>
import SearchTextField from '@/components/atoms/SearchTextField.vue';
import { useAppStore } from '@/stores/app';

type Props = {
  search: string;
};
type Emit = {
  (e: 'update:search', search: string): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>
<template>
  <search-text-field
    :search="props.search"
    @update:search="emits('update:search', $event)"
    :style="`
      width: 200px;
      background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
    `"
  />
</template>