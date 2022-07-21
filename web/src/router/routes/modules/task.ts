import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const dashboard: AppRouteModule = {
  path: '/task',
  name: 'Task',
  component: LAYOUT,
  redirect: '/task/index',
  meta: {
    orderNo: 10,
    icon: 'ant-design:clock-circle-filled',
    title: t('routes.task.tasks'),
  },
  children: [
    {
      path: 'index',
      name: 'Index',
      component: () => import('/@/views/dashboard/analysis/index.vue'),
      meta: {
        // affix: true,
        title: t('routes.task.index'),
      },
    },
    {
      path: 'logs',
      name: 'Logs',
      component: () => import('/@/views/dashboard/workbench/index.vue'),
      meta: {
        title: t('routes.task.logs'),
      },
    },
  ],
};

export default dashboard;
