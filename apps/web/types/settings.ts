import { Base } from "./base";

export interface UserSettings {
  id: string;
  user_id: string;
  notify_dashboard: boolean;
  notify_email: boolean;
  notify_webhook: boolean;
  webhook_url?: string;
  email?: string;
  smtp_host?: string;
  smtp_port?: number;
  smtp_username?: string;
  smtp_password?: string;
}

export type UpdateSettingsRequest = Partial<Omit<UserSettings, 'id' | 'user_id'>>;

export type SettingsResponse = Base<UserSettings>;
export type GetSettingsResponse = SettingsResponse;
export type UpdateSettingsResponse = SettingsResponse;
