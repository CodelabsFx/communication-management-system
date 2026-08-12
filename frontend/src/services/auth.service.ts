import api from './api';

export interface LoginPayload {
  email: string;
  password: string;
}

export interface AuthResponse {
  token: string;
}

export async function login(payload: LoginPayload) {
  const response = await api.post<AuthResponse>('/login', payload);
  return response.data;
}

export function logout() {
  localStorage.removeItem('auth_token');
}
