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
}

export interface ConnectionStatus {
  isConnected: boolean;
  lastSync: string;
  error?: string;
}

export interface ConnectionStats {
  total_connections: number;
  total_size: number;
  average_size: number;
  ssl_percentage: number;
};
