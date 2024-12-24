export interface BackupConfig {
  id: string;
  connectionId: string;
  schedule: string; // Cron expression
  retentionDays: number;
  lastBackup: string;
  status: 'pending' | 'running' | 'completed' | 'failed';
  backupPath: string;
}

export interface BackupStatus {
  id: string;
  status: BackupConfig['status'];
  progress?: number;
  message?: string;
}
