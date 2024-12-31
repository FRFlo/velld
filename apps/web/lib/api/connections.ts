import { Connection, ConnectionStatus, ConnectionStats, BaseConnection } from '@/types/connection';
import { apiRequest } from '@/lib/api-client';

export async function testConnection(connection: BaseConnection): Promise<ConnectionStatus> {
  return apiRequest<ConnectionStatus>('/api/connections/test', {
    method: 'POST',
    body: JSON.stringify(connection),
  });
}

export async function saveConnection(connection: BaseConnection): Promise<void> {
  return apiRequest('/api/connections', {
    method: 'POST',
    body: JSON.stringify(connection),
  });
}

export async function getConnections(): Promise<Connection[]> {
  return apiRequest<Connection[]>('/api/connections');
}

export async function getConnectionStats(): Promise<ConnectionStats> {
  return apiRequest<ConnectionStats>('/api/connections/stats');
}