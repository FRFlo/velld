import { BackupListResponse } from '@/types/backup';
import { apiRequest } from '../api-client';

export interface GetBackupsParams {
  page?: number;
  limit?: number;
  search?: string;
}

export async function saveBackup(connectionId: string): Promise<void> {
  return apiRequest('/api/backups', {
    method: 'POST',
    body: JSON.stringify({ connection_id: connectionId }),
  });
}

export async function getBackups(params?: GetBackupsParams): Promise<BackupListResponse> {
  let url = `/api/backups`;

  if (params?.page && params?.limit) {
    url += `?page=${params.page}&limit=${params.limit}`;
  }

  if (params?.search) {
    url += `&search=${params.search}`;
  }

  return apiRequest<BackupListResponse>(url, {
    method: 'GET',
  });
}