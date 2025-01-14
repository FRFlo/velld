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
  const searchParams = new URLSearchParams();
  if (params?.page) searchParams.append('page', params.page.toString());
  if (params?.limit) searchParams.append('limit', params.limit.toString());
  if (params?.search) searchParams.append('search', params.search);

  const queryString = searchParams.toString();
  const url = `/api/backups${queryString ? `?${queryString}` : ''}`;

  return apiRequest<BackupListResponse>(url, {
    method: 'GET',
  });
}