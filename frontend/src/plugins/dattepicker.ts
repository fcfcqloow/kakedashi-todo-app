import { App, Plugin } from 'vue';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css';

export const datePickerPlugin: Plugin = {
    install: (app: App, options?: any) => {
      app.component('VueDatePicker', VueDatePicker);
    },
  };

  export default datePickerPlugin;
