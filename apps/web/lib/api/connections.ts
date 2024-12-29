import { DatabaseConnection, ConnectionStatus, ConnectionStats } from '@/types/connection';
import { apiRequest } from '@/lib/api-client';

export async function testConnection(connection: DatabaseConnection): Promise<ConnectionStatus> {
  return apiRequest<ConnectionStatus>('/api/connections/test', {
    method: 'POST',
    body: JSON.stringify(connection),
  });
}

export async function saveConnection(connection: DatabaseConnection): Promise<void> {
  return apiRequest('/api/connections', {
    method: 'POST',
    body: JSON.stringify(connection),
  });
}

export async function getConnections(): Promise<DatabaseConnection[]> {
  return apiRequest<DatabaseConnection[]>('/api/connections');
}

export async function getConnectionStats(): Promise<ConnectionStats> {
  return apiRequest<ConnectionStats>('/api/connections/stats');
}