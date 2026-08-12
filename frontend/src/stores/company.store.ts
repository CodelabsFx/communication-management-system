import { defineStore } from 'pinia';

interface CompanyState {
  id: string | null;
  name: string | null;
}

export const useCompanyStore = defineStore('company', {
  state: (): CompanyState => ({
    id: null,
    name: null,
  }),
  actions: {
    setCompany(id: string, name: string) {
      this.id = id;
      this.name = name;
    },
    clearCompany() {
      this.id = null;
      this.name = null;
    },
  },
});
