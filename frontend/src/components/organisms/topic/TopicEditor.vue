<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Topic } from '@/core/Topic';
import TopicIcons from './TopicIcons.vue';
import { useAppStore } from '../../../stores/app';
import TopicListTable from './TopicListTable.vue';

type Props = {
  topics: Array<Topic>;
};
type Emit = {
  (e: 'update:topic', topic: Topic): void;
  (e: 'create:topic', topic: Topic): void;
  (e: 'remove:topic', topic: Topic): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const newTopic = ref(new Topic({ value : '',  created : new Date(), color :  '#' + Math.floor(Math.random() * 0xFFFFFF).toString(16) }));
const showDeleteIcon = ref(false);
const isValid = computed(() => !newTopic.value.value || !newTopic.value.color);
const onCreate = () => {
  if (newTopic.value.value) {
    newTopic.value.created = new Date();
    emits('create:topic', newTopic.value);
    newTopic.value = new Topic({ value : '',  created : new Date(), color : '#'+Math.floor(Math.random() * 0xFFFFFF).toString(16) });
  }
};

const shortcut = (event: KeyboardEvent) => {
  if (event.shiftKey) {
    switch(event.code) {
      case 'Enter':
        onCreate();
        return;
    }
  }
};

</script>

<template>
  <v-row justify="center" style="margin: 0;">
    <v-col cols="12" align-self="center">
      <topic-icons
        v-if="settings.topicListType === 'icon'"
        height="50vh"
        width="100%"
        v-model:show-delete-icon="showDeleteIcon"
        :topics="props.topics?.filter((v) => !v.done) ?? []"
        @delete="emits('remove:topic', $event)"
      />
      <topic-list-table
        v-else
        :topics="props.topics"
        height="50vh"
        width="100%"
        @remove="emits('remove:topic', $event)"
      />
    </v-col>
    <v-col>
      <v-text-field v-model="newTopic.value" label="value" @keydown="shortcut" />
      <v-color-picker v-model="newTopic.color" hide-inputs hide-sliders show-swatches ></v-color-picker>
    </v-col>
    <v-col>
      <v-btn
        :disabled="isValid"
        :color="isValid ? 'black' : 'green'"
        icon="mdi-check"
        @click="onCreate"
      />
    </v-col>
    <v-col>
      <v-sheet
        style="margin: 0; overflow: scroll;"
        width="100%"
        height="100%"
      >
      <h4 class="text-h5 font-weight-bold mb-4">Done Area</h4>
      </v-sheet>
    </v-col>
  </v-row>
</template>

  <style scoped>
.border-radius {
  width: 50px;
  height: 50px;
  line-height: 50px;
  background-color: green;
  border-radius: 50%;
  color: #fff;
  text-align: center;
  display: inline-block; _display: inline;
  margin: 10px;
  cursor: pointer;
  overflow: hidden;
}
.border-radius p {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
