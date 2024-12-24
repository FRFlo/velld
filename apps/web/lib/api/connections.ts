import { DatabaseConnection, ConnectionStatus } from '@/types/connection';

const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function testConnection(connection: DatabaseConnection): Promise<ConnectionStatus> {
  const response = await fetch(`${API_BASE}/api/connections/test`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(connection),
  });

  if (!response.ok) {
    throw new Error('Failed to test connection');
  }

  return response.json();
}

export async function saveConnection(connection: DatabaseConnection): Promise<void> {
  const response = await fetch(`${API_BASE}/api/connections`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(connection),
  });

  if (!response.ok) {
    throw new Error('Failed to save connection');
  }
}
