import api from './api';

export interface MessageItem {
  id: string;
  conversationId: string;
  content: string;
  senderId: string;
  createdAt: string;
}

export async function fetchMessages(conversationId: string) {
  const response = await api.get<MessageItem[]>(`/conversations/${conversationId}/messages`);
  return response.data;
}
