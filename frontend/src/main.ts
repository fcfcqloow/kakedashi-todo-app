import { createApp, type Plugin } from 'vue';

import App from './App.vue';

const app = createApp(App);
Object.values(import.meta.glob<Plugin>('./plugins/*.ts', { eager: true, import: 'default' }))
  .forEach(app.use);

app.mount('#app');
