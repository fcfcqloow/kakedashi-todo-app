import { createRouter, createWebHashHistory } from 'vue-router/auto';
import { setupLayouts } from 'virtual:generated-layouts';

export default createRouter({
  history: createWebHashHistory(''),
  extendRoutes: setupLayouts,
})
