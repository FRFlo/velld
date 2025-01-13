export interface BackupList {
  id: string;
  connetion_id: string;
  database_type: string;
  size: number;
  status: string;
  path: string;
  scheduled_time: string;
  started_time: string;
  completed_time: string;
  created_at: string;
  updated_at: string;
}