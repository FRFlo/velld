'use client';

import { Clock, Calendar } from 'lucide-react';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Switch } from "@/components/ui/switch";
import { Label } from '@/components/ui/label';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import type { Connection } from '@/types/connection';
import { useBackup } from '@/hooks/use-backup';
import { useState, useEffect } from 'react';

const CRON_SCHEDULES = {
  'hourly': '0 0 * * * *',
  'daily': '0 0 0 * * *',
  'weekly': '0 0 0 * * 0',
  'monthly': '0 0 0 1 * *'
};

const RETENTION_DAYS = {
  '7': 7,
  '30': 30,
  '90': 90,
  '365': 365
};

interface BackupScheduleDialogProps {
  connectionId: string | null;
  connection?: Connection;
  isEnabled?: boolean;
  onClose: () => void;
}

export function BackupScheduleDialog({
  connectionId,
  connection,
  isEnabled,
  onClose,
}: BackupScheduleDialogProps) {
  const { createSchedule, updateExistingSchedule, disableSchedule, isScheduling, isDisabling, isUpdating } = useBackup();
  const [enabled, setEnabled] = useState(isEnabled);

  useEffect(() => {
    setEnabled(isEnabled);
  }, [isEnabled]);

  if (!connectionId || !connection) return null;

  const handleScheduleChange = async (checked: boolean) => {
    try {
      if (checked) {
        await createSchedule({
          connection_id: connectionId,
          cron_schedule: CRON_SCHEDULES['daily'],
          retention_days: RETENTION_DAYS['30']
        });
        setEnabled(true);
      } else {
        await disableSchedule(connectionId);
        setEnabled(false);
      }
    } catch (error) {
      // If there's an error, revert the switch state
      setEnabled(!checked);
      console.error('Failed to update schedule:', error);
    }
  };

  const handleScheduleSubmit = async (schedule: string, retention: string) => {
    try {
      if (enabled) {
        // Update existing schedule
        await updateExistingSchedule({
          connectionId,
          params: {
            cron_schedule: CRON_SCHEDULES[schedule as keyof typeof CRON_SCHEDULES],
            retention_days: RETENTION_DAYS[retention as keyof typeof RETENTION_DAYS]
          }
        });
      } else {
        // Create new schedule
        await createSchedule({
          connection_id: connectionId,
          cron_schedule: CRON_SCHEDULES[schedule as keyof typeof CRON_SCHEDULES],
          retention_days: RETENTION_DAYS[retention as keyof typeof RETENTION_DAYS]
        });
        setEnabled(true);
      }
    } catch (error) {
      console.error('Failed to update schedule:', error);
    }
  };

  return (
    <Dialog open={!!connectionId} onOpenChange={() => onClose()}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Backup Schedule - {connection.name}</DialogTitle>
        </DialogHeader>
        <div className="space-y-6 py-4">
          <div className="flex items-center justify-between space-x-4">
            <div className="space-y-0.5">
              <Label>Automatic Backups</Label>
              <div className="text-sm text-muted-foreground">
                Enable scheduled backups for this database
              </div>
            </div>
            <Switch
              checked={enabled}
              disabled={isScheduling || isDisabling || isUpdating}
              onCheckedChange={handleScheduleChange}
            />
          </div>

          {enabled && (
            <>
              <div className="space-y-2">
                <Label className="text-sm flex items-center">
                  <Clock className="h-4 w-4 mr-1" />
                  Backup Frequency
                </Label>
                <Select
                  defaultValue="daily"
                  onValueChange={(value) => handleScheduleSubmit(value, '30')}
                  disabled={isScheduling}
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="hourly">Every Hour</SelectItem>
                    <SelectItem value="daily">Daily</SelectItem>
                    <SelectItem value="weekly">Weekly</SelectItem>
                    <SelectItem value="monthly">Monthly</SelectItem>
                  </SelectContent>
                </Select>
              </div>

              <div className="space-y-2">
                <Label className="text-sm flex items-center">
                  <Calendar className="h-4 w-4 mr-1" />
                  Retention Period
                </Label>
                <Select
                  defaultValue="30"
                  onValueChange={(value) => handleScheduleSubmit('daily', value)}
                  disabled={isScheduling}
                >
                  <SelectTrigger>
                    <SelectValue />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="7">7 Days</SelectItem>
                    <SelectItem value="30">30 Days</SelectItem>
                    <SelectItem value="90">90 Days</SelectItem>
                    <SelectItem value="365">1 Year</SelectItem>
                  </SelectContent>
                </Select>
              </div>
            </>
          )}
        </div>
      </DialogContent>
    </Dialog>
  );
} 