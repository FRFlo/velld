export interface DatabaseConnection {
  id: string;
  name: string;
  type: 'mysql' | 'postgresql' | 'mongodb';
  host: string;
  port: number;
  status: 'connected' | 'disconnected';
  username: string;
  password: string;
  database: string;
  ssl: boolean;
  last_connected_at: string;
  database_size: number;
}

export interface ConnectionStatus {
  isConnected: boolean;
  lastSync: string;
  error?: string;
}
