<script lang="ts" setup>
import { ref } from 'vue';
import { useAppStore } from '@/stores/app';
import TopicSelector from '@/components/organisms/topic/TopicSelector.vue';
import { Topic } from '@/core/Topic';
import { Color } from '@/core/Color';

type Props = {
  topics?: Topic[];
  filterTopics: Topic[];
};
type Emit = {
  (e: 'update:filterTopics', topics: Topic[]): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();
const showField = ref(false);

</script>
<template>
  <topic-selector
    v-if="showField"
    :selected="props.filterTopics"
    @update:selected="emits('update:filterTopics', $event)"
    placeholder="Topic Filter"
    :topics="props.topics"
    :style="`
      background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
      width: 50px;
      height: 56px;
      overflow: scroll;
      margin: 10px;
    `"
  />
  <v-tooltip
    text="Filter Topics"
    location="bottom"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        v-bind="props"
        size="small"
        variant="outlined"
        :icon="showField ? 'mdi-close-circle' : 'mdi-filter'"
        :color="showField ? 'red' : ''"
        :style="`
          margin: 10px;
          background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
        `"
        @click="showField = !showField"
      />
    </template>
  </v-tooltip>
</template>