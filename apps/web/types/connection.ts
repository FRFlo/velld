export interface BaseConnection {
  name: string;
  type: 'mysql' | 'postgresql' | 'mongodb';
  host: string;
  port: number;
  username: string;
  password: string;
  database: string;
  ssl: boolean;
}

export interface Connection extends BaseConnection {
  id: string;
  status: 'connected' | 'disconnected';
  last_connected_at: string;
  database_size: number;
  last_backup_time: string;
  backup_enabled: boolean;
  cron_schedule?: string;
  retention_days?: number;
}

export interface ConnectionStatus {
  isConnected: boolean;
  lastSync: string;
  error?: string;
}