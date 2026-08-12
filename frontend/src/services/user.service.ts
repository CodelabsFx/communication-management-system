import api from './api';
import type { UserProfile } from '../types';

export async function fetchCurrentUser() {
  const response = await api.get<UserProfile>('/user/me');
  return response.data;
}
