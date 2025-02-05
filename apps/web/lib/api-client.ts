export const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = 'ApiError';
  }
}

interface ExtendedRequestInit extends RequestInit {
  responseType?: 'json' | 'blob';
}

export async function apiRequest<T>(
  endpoint: string,
  options: ExtendedRequestInit = {}
): Promise<T> {
  const token = localStorage.getItem('token');
  const headers = {
    'Content-Type': 'application/json',
    ...(token && { Authorization: `Bearer ${token}` }),
    ...options.headers,
  };

  const response = await fetch(`${API_BASE}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    if (response.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    const error = await response.text();
    throw new ApiError(response.status, error || `HTTP error! status: ${response.status}`);
  }

  return options.responseType === 'blob' ? response.blob() : response.json();
}
