import { useState } from 'react';
import { useToast } from '@/hooks/use-toast';
import type { BackupConfig } from './types';

export function useBackupConfigs() {
  const { toast } = useToast();
  const [backupConfigs, setBackupConfigs] = useState<Record<string, BackupConfig>>({});

  const handleBackupNow = (connectionId: string) => {
    toast({
      title: "Backup started",
      description: "Your database backup has been initiated.",
    });
    setBackupConfigs(prev => ({
      ...prev,
      [connectionId]: {
        ...prev[connectionId],
        lastBackup: new Date().toISOString(),
      }
    }));
  };

  const handleUpdateBackupConfig = (connectionId: string, config: Partial<BackupConfig>) => {
    setBackupConfigs(prev => ({
      ...prev,
      [connectionId]: {
        enabled: config.enabled ?? prev[connectionId]?.enabled ?? false,
        schedule: config.schedule || prev[connectionId]?.schedule || 'daily',
        retention: config.retention || prev[connectionId]?.retention || '30',
        lastBackup: prev[connectionId]?.lastBackup,
      }
    }));

    toast({
      title: config.enabled !== undefined ? "Backup schedule updated" : "Settings saved",
      description: config.enabled 
        ? config.enabled 
          ? "Automatic backups have been enabled" 
          : "Automatic backups have been disabled"
        : "Your backup configuration has been saved.",
    });
  };

  return {
    backupConfigs,
    handleBackupNow,
    handleUpdateBackupConfig,
  };
} 