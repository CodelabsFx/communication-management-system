import { defineStore } from 'pinia';

interface AuthState {
  token: string | null;
  userId: string | null;
  companyId: string | null;
  permissions: string[];
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    token: null,
    userId: null,
    companyId: null,
    permissions: []
  }),
  getters: {
    isAuthenticated: (state) => !!state.token
  },
  actions: {
    setSession(token: string, userId: string, companyId: string, permissions: string[]) {
      this.token = token;
      this.userId = userId;
      this.companyId = companyId;
      this.permissions = permissions;
    },
    clearSession() {
      this.token = null;
      this.userId = null;
      this.companyId = null;
      this.permissions = [];
    }
  }
});
