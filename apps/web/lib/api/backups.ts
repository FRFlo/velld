import { BackupList } from '@/types/backup';
import { apiRequest } from '../api-client';

export async function saveBackup(connectionId: string): Promise<void> {
  return apiRequest('/api/backups', {
    method: 'POST',
    body: JSON.stringify({ connection_id: connectionId }),
  });
}

export async function getBackups(): Promise<BackupList[]> {
  return apiRequest('/api/backups', {
    method: 'GET',
  });
}