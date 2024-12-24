import { BackupConfig, BackupStatus } from '@/types/backup';

const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function createBackupSchedule(config: BackupConfig): Promise<void> {
  const response = await fetch(`${API_BASE}/api/backups/schedule`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(config),
  });

  if (!response.ok) {
    throw new Error('Failed to create backup schedule');
  }
}

export async function getBackupStatus(backupId: string): Promise<BackupStatus> {
  const response = await fetch(`${API_BASE}/api/backups/${backupId}/status`);

  if (!response.ok) {
    throw new Error('Failed to get backup status');
  }

  return response.json();
}
