import { apiRequest } from "@/lib/api-client";
import { UpdateSettingsRequest, GetSettingsResponse, UpdateSettingsResponse } from "@/types/settings";

export async function getUserSettings(): Promise<GetSettingsResponse> {
  return apiRequest<GetSettingsResponse>('/api/settings');
}

export async function updateUserSettings(settings: UpdateSettingsRequest): Promise<UpdateSettingsResponse> {
  return apiRequest<UpdateSettingsResponse>('/api/settings', {
    method: 'PUT',
    body: JSON.stringify(settings),
  });
}
