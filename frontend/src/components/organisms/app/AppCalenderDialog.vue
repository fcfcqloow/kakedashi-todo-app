<script lang="ts" setup>
import Calender from '@/components/atoms/Calender.vue';
import TimeSelector from '@/components/atoms/TimeSelector.vue';
import { Color } from '@/core/Color';
import { useAppStore } from '@/stores/app';
import { onMounted, ref, computed, watch } from 'vue';
import { getFontColor } from '../../../core/Color';

const { settings, calenderDialog, closeCalender } = useAppStore();

const time = ref(`${calenderDialog.initDate.getHours().toString().padStart(2, '0')}:${calenderDialog.initDate.getMinutes().toString().padStart(2, '0')}`);
const date = ref(calenderDialog.initDate);
const width = 250;
const height = 300;
const x = ref(0);
const y = ref(0);
const isShow = ref(false);
const viewportWidth = ref(window.innerWidth);
const viewportHeight = ref(window.innerHeight);
const isLeft = computed(() => (viewportWidth.value - x.value - width) > 0);
const isTop = computed(() => (viewportHeight.value - y.value - height) > 0);

onMounted(() => {
  window.addEventListener('resize', () => {
    viewportWidth.value = window.innerWidth;
    viewportHeight.value = window.innerHeight;
  });

  document.body.addEventListener( "click", (event) => {
    if(isShow.value && !(event?.target as any)?.closest('#calender')) {
      closeCalender(); 
    }

    if (!isShow.value) {
      x.value = event.pageX;
	    y.value = event.pageY;
    }
  });
});

watch(() => calenderDialog.initDate, () => {
  date.value = calenderDialog.initDate;
  time.value = `${calenderDialog.initDate.getHours().toString().padStart(2, '0')}:${calenderDialog.initDate.getMinutes().toString().padStart(2, '0')}`;
});

watch(() => calenderDialog.isShow, () => {
  if (calenderDialog.isShow) {
    setTimeout(() => {
      calenderDialog.isShow && (isShow.value = true);
    }, 100);
  } else {
    isShow.value = false; 
  }
});

const applyDate = () => {
  const [hours, seconds] = time.value.split(':');
  date.value.setHours(Number(hours));
  date.value.setMinutes(Number(seconds));
  calenderDialog?.dateCallback && calenderDialog.dateCallback(date.value);
  closeCalender();
};

</script>

<template>
  <div>
    <calender
      v-if="calenderDialog.isShow"
      v-model="date"
      id="calender"
      :style="`
        position: fixed;
        z-index: 999;
        left: ${isLeft ? x  : x-width}px;
        top: ${isTop ? y : y-height}px;
      `"
      :width="width"
      :height="height"
      :color="settings.mode === 'dark' ? Color.BLACK : Color.WHITE"
    >
      <template #bottom>
        <div :style="`background-color: ${settings.mode === 'dark' ? Color.WHITE : Color.BLACK}`">
          <time-selector v-model="time" style="width: 100%;" />
        </div>
        <v-sheet :color="settings.mode === 'dark' ? Color.BLACK : Color.WHITE">
          <v-card-actions>
            <span :style="`
              color: ${getFontColor(settings.mode === 'dark' ? Color.BLACK : Color.WHITE)}
            `">
              {{ date.toLocaleDateString() }} {{ time }}
            </span>
          <v-spacer></v-spacer>
          <v-btn color="blue-darken-1" variant="text" @click="applyDate">OK</v-btn>
        </v-card-actions>
        </v-sheet>
      </template>
    </calender>
  </div>
</template>