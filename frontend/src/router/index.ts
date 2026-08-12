import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import HomeView from '../views/HomeView.vue';
import DashboardPage from '../features/dashboard/DashboardPage.vue';
import MessagesPage from '../features/communication/messages/MessagesPage.vue';

const routes = [
  { path: '/', name: 'home', component: HomeView },
  { path: '/login', name: 'login', component: LoginView },
  { path: '/dashboard', name: 'dashboard', component: DashboardPage },
  { path: '/communication/messages', name: 'messages', component: MessagesPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
