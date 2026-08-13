import { createRouter, createWebHistory } from 'vue-router';
import LoginView from '../views/LoginView.vue';
import HomeView from '../views/HomeView.vue';
import RegisterView from '../views/RegisterView.vue';
import ProfileView from '../views/ProfileView.vue';

// Dashboard
import DashboardPage from '../features/dashboard/DashboardPage.vue';

// Communication
import MessagesPage from '../features/communication/messages/MessagesPage.vue';
import ChannelsPage from '../features/communication/channels/ChannelsPage.vue';
import AnnouncementsPage from '../features/communication/announcements/AnnouncementsPage.vue';

// Work: tasks, issues, projects
import TasksPage from '../features/work/tasks/TasksPage.vue';
import TaskDetailPage from '../features/work/tasks/TaskDetailPage.vue';
import IssuesPage from '../features/work/issues/IssuesPage.vue';
import IssueDetailPage from '../features/work/issues/IssueDetailPage.vue';
import ProjectsPage from '../features/work/projects/ProjectsPage.vue';
import ProjectDetailPage from '../features/work/projects/ProjectDetailPage.vue';

// Meetings
import MeetingsPage from '../features/meetings/MeetingsPage.vue';
import MeetingDetailPage from '../features/meetings/MeetingDetailPage.vue';

// Documents
import DocumentsPage from '../features/documents/DocumentsPage.vue';

// Employees
import EmployeesPage from '../features/employees/EmployeesPage.vue';
import EmployeeDetailPage from '../features/employees/EmployeeDetailPage.vue';

// Notifications
import NotificationsPage from '../features/notifications/NotificationsPage.vue';

// Reports
import ReportsPage from '../features/reports/ReportsPage.vue';

// Administration
import UsersAdminPage from '../features/administration/UsersAdminPage.vue';
import DepartmentsAdminPage from '../features/administration/DepartmentsAdminPage.vue';
import TeamsAdminPage from '../features/administration/TeamsAdminPage.vue';
import RolesAdminPage from '../features/administration/RolesAdminPage.vue';
import PermissionsAdminPage from '../features/administration/PermissionsAdminPage.vue';
import AuditLogsPage from '../features/administration/AuditLogsPage.vue';

// Settings
import SettingsPage from '../features/settings/SettingsPage.vue';

const routes = [
  { path: '/', name: 'home', component: HomeView },
  { path: '/login', name: 'login', component: LoginView },
  { path: '/register', name: 'register', component: RegisterView },
  { path: '/profile', name: 'profile', component: ProfileView },

  { path: '/dashboard', name: 'dashboard', component: DashboardPage },

  // Communication
  { path: '/communication/messages', name: 'messages', component: MessagesPage },
  { path: '/communication/channels', name: 'channels', component: ChannelsPage },
  { path: '/communication/announcements', name: 'announcements', component: AnnouncementsPage },

  // Work
  { path: '/work/tasks', name: 'tasks', component: TasksPage },
  { path: '/work/tasks/:id', name: 'task-detail', component: TaskDetailPage, props: true },
  { path: '/work/issues', name: 'issues', component: IssuesPage },
  { path: '/work/issues/:id', name: 'issue-detail', component: IssueDetailPage, props: true },
  { path: '/work/projects', name: 'projects', component: ProjectsPage },
  { path: '/work/projects/:id', name: 'project-detail', component: ProjectDetailPage, props: true },

  // Meetings
  { path: '/meetings', name: 'meetings', component: MeetingsPage },
  { path: '/meetings/:id', name: 'meeting-detail', component: MeetingDetailPage, props: true },

  // Documents
  { path: '/documents', name: 'documents', component: DocumentsPage },

  // Employees
  { path: '/employees', name: 'employees', component: EmployeesPage },
  { path: '/employees/:id', name: 'employee-detail', component: EmployeeDetailPage, props: true },

  // Notifications
  { path: '/notifications', name: 'notifications', component: NotificationsPage },

  // Reports
  { path: '/reports', name: 'reports', component: ReportsPage },

  // Administration
  { path: '/admin/users', name: 'admin-users', component: UsersAdminPage },
  { path: '/admin/departments', name: 'admin-departments', component: DepartmentsAdminPage },
  { path: '/admin/teams', name: 'admin-teams', component: TeamsAdminPage },
  { path: '/admin/roles', name: 'admin-roles', component: RolesAdminPage },
  { path: '/admin/permissions', name: 'admin-permissions', component: PermissionsAdminPage },
  { path: '/admin/audit-logs', name: 'admin-audit-logs', component: AuditLogsPage },

  // Settings
  { path: '/settings', name: 'settings', component: SettingsPage }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
