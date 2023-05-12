<script lang="ts" setup>
import { computed, ref, watch } from 'vue';
import { getEndDay, getStartWeek, nextMonth, backMonth } from '@/util/date';
import { Color, getFontColor } from '@/core/Color';

type Props = {
  modelValue?: Date;
  width?: string|number;
  height?: string|number;
  todyColor?: string;
  color?: string;
  pickColor?: string;
};

type Emit = {
  (e: 'update:modelValue', date: Date): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const isHoverTodyBtn = ref(false);
const today = new Date();
const pickingDate = ref<Date|null>(new Date(props.modelValue ?? today));
const date = ref(new Date(props.modelValue ?? today));
const year = computed(() => date.value.getFullYear());
const month = computed(() => date.value.getMonth() + 1);

const calendervalue = computed(() => {
  const next = nextMonth(date.value);
  const afterYear = next.getFullYear();
  const afterMonth = next.getMonth() + 1;

  const back = backMonth(date.value);
  const beforeYear = back.getFullYear();
  const beforeMonth = back.getMonth() + 1;

  const endDay = getEndDay(year.value, month.value);
  const startWeek = getStartWeek(year.value, month.value);
  let targetDay = 1;
  let nextMonthDay = 1;
  return new Array(6).fill(null).map((_, idx) => {
    return new Array(7).fill(null).map((_, week) => {
      if (idx === 0) {
        return (startWeek <= week)
          ? { year : year.value, month : month.value, day : targetDay++ }
          : { year : beforeYear, month : beforeMonth, day : getEndDay(beforeYear, beforeMonth) - startWeek + (week + 1)};
      }

      return idx*7+week - startWeek + 1 <= endDay
        ? { year : year.value, month : month.value, day : targetDay++ }
        : { year : afterYear, month : afterMonth, day : nextMonthDay++ };
    });
  });
});

const clickDate = (date: { day: number; month: number; year: number; }) => {
  const newDate = new Date(`${date.year}/${date.month}/${date.day}`);
  emits('update:modelValue', newDate);
  pickingDate.value = newDate;
};

watch(()=> props.modelValue, () => {
  if (props.modelValue) {
    pickingDate.value = props.modelValue;
  }
});

</script>

<template>
  <div>
    <v-card
      :style="`min-height: 100px; color: ${props.color && getFontColor(props.color)}`"
      :width="props.width ?? '300px'"
      :color="props.color"
    >
      <v-btn 
        position="absolute"
        @mouseover="isHoverTodyBtn=true"
        @mouseleave="isHoverTodyBtn=false"
        :color="props.todyColor"
        @click="date = new Date()"
        size="10"
      >
        <v-icon v-if="!isHoverTodyBtn" >mdi-calendar-today</v-icon>
        <div style="font-size: 3px;" v-else>Today</div>
      </v-btn>

      <v-card-title style="text-align: center;">
        <v-btn size="30" icon="mdi-arrow-left" @click="date = new Date(date.getFullYear(), date.getMonth() - 1)"/>
        {{ `${year}/${month}` }}
        <v-btn size="30" icon="mdi-arrow-right" @click="date = new Date(date.getFullYear(), date.getMonth() + 1)" />
      </v-card-title>

      <v-row v-for="week in calendervalue" no-gutters :style="`height: calc(${100/calendervalue.length}% - ${60/calendervalue.length}px)`">
        <v-col v-for="(day, weekIndex) in week" :key="JSON.stringify(day)" :style="`
          width: ${100/7}%;
          border: solid;
          border-color: rgba(155, 155, 155, .6);
        `" >
          <v-btn
            @click="clickDate(day);"
            :color="pickingDate?.getFullYear() === day.year && pickingDate.getMonth()+1 === day.month && pickingDate.getDate() === day.day ?  Color.GRAY : undefined"
            :style="`
              height: 100%;
              width: 100%;
              min-width: 0px;
              padding: 0;
              cursor: pointer;
              white-space: nowrap;
              opacity: ${day.month === month ? '1.0' : '0.3'};
              color: ${weekIndex === 0 && 'red' || weekIndex === 6 && 'blue' || undefined};
              text-align: center;
            `"
          >
            <div
              :style="`
                background-color: ${
                  today.getFullYear() === day.year
                  && today.getMonth() + 1 === day.month
                  && today.getDate() === day.day
                    ? (props.todyColor ?? 'red')
                    : ''
                };
                border-radius: 50%;
              `"
            >
              {{ day.day }}
            </div>
          </v-btn>
        </v-col>
      </v-row>
    </v-card>
    <slot name="bottom" />
  </div>
 
</template>
