import { Auth, AuthResponse, Profile } from "@/types/auth";

const API_BASE = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function login(auth: Auth): Promise<AuthResponse> {
  const response = await fetch(`${API_BASE}/api/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(auth),
  });

  if (!response.ok) {
    throw new Error('Failed to login');
  }

  const data = await response.json();
  localStorage.setItem('token', data.token);
  return data;
}

export async function register(auth: Auth): Promise<void> {
  const response = await fetch(`${API_BASE}/api/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(auth),
  });
    
  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || 'Failed to register');
  }
}

export async function getProfile(): Promise<Profile> {
  const token = localStorage.getItem('token');
  const response = await fetch(`${API_BASE}/api/auth/profile`, {
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Failed to get profile');
  }

  return response.json();
}