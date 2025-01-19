export type SortBy = 'name' | 'status' | 'type' | 'lastBackup';

export interface BackupConfig {
  enabled: boolean;
  schedule: string;
  retention: string;
  lastBackup?: string;
}

export const statusColors = {
  connected: 'bg-emerald-500/15 text-emerald-500 border-emerald-500/20',
  disconnected: 'bg-amber-500/15 text-amber-500 border-amber-500/20',
  error: 'bg-red-500/15 text-red-500 border-red-500/20',
} as const;

export const typeLabels = {
  mysql: 'MySQL',
  postgresql: 'PostgreSQL',
  mongodb: 'MongoDB',
} as const; 