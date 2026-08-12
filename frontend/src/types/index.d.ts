export interface UserProfile {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  companyId: string;
  permissions: string[];
}

export interface CompanyInfo {
  id: string;
  name: string;
}
