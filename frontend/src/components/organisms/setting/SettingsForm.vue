<script lang="ts" setup>
import { ref } from 'vue';
import { Settings } from '@/core/Settings';
import RadioButtons from '@/components/atoms/RadioButtons.vue';
import { Color } from '@/core/Color';

type Props = {
  settings: Settings;
};
type Emit = {
  (e: 'submit', settings: Settings): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const mode = ref<'dark'|'white'>(props.settings.mode);
const limit = ref(props.settings.taskLimit);
const sortable = ref(String(props.settings.sortable));
const topiccolor = ref(props.settings.topicColorFlag ? 'on' : 'off');

const alignment = ref(props.settings.topicListType === 'list' ? 0 : 1);

const interval = ref(props.settings.notificationIntervalSec);
</script>

<template>
  <form>
    <v-row>
      <v-col cols="12">
        <v-card-title>Settings</v-card-title>
        <v-divider></v-divider>
      </v-col>
    </v-row>

    <v-row style="margin-top: 20px;">
      <v-col cols="12">
        <v-card-subtitle>TODO List</v-card-subtitle>
        <v-divider></v-divider>
      </v-col>
      <v-col cols="12">
        <v-row>
          <v-col cols="2">Theme</v-col>
          <v-col>
            <radio-buttons
            v-model="mode"
            :labels="['while', 'dark']"
            inline
            />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="2">Get Limit</v-col>
          <v-col><v-text-field v-model="limit" /></v-col>
        </v-row>

        <v-row>
          <v-col cols="2">Sortable</v-col>
          <v-col><v-select v-model="sortable" :items="['true', 'false']" label="Sortable" /></v-col>
        </v-row>

        <v-row>
          <v-col cols="2">Topic Color</v-col>
          <v-col><v-select v-model="topiccolor" :items="['on', 'off']" label="Topic Color Flag" /></v-col>
        </v-row>
      </v-col>
    </v-row>

    <v-row style="margin-top: 20px;">
      <v-col cols="12">
        <v-card-subtitle>Topic Editor</v-card-subtitle>
        <v-divider></v-divider>
      </v-col>
      <v-col cols="12">
        <v-row>
          <v-col cols="2">List Type</v-col>
          <v-col>
            <v-btn-toggle
              v-model="alignment"
              variant="outlined"
              divided
            >
              <v-btn>
                <v-icon icon="mdi-format-list-bulleted"></v-icon>
              </v-btn>

              <v-btn>
                <v-icon icon="mdi-dots-circle"></v-icon>
              </v-btn>

            </v-btn-toggle>
          </v-col>
        </v-row>
      </v-col>
    </v-row>

    <v-row style="margin-top: 20px;">
      <v-col cols="12">
        <v-card-subtitle>Notification</v-card-subtitle>
        <v-divider></v-divider>
      </v-col>
      <v-col cols="12">

        <v-row>
          <v-col cols="2">Notification Interval (sec: min30s)</v-col>
          <v-col><v-text-field
            :model-value="interval"
            @update:model-value="interval = Math.max(30, Number($event))"
          /></v-col>
        </v-row>

      </v-col>
    </v-row>

    <div class="text-end">
      <v-btn
        class=""
        :color="Color.SUCCESS"
        style="margin-top: 50px;"
        @click="emits(
          'submit',
          {
            ...settings,
            mode                    : mode,
            sortable                : sortable === 'true',
            topicColorFlag          : topiccolor === 'on',
            topicListType           : alignment === 0 ? 'list' : 'icon',
            taskLimit               : Number(limit),
            notificationIntervalSec : interval,
          },
        )"
      >
        Submit
      </v-btn>
    </div>
  </form>
</template>