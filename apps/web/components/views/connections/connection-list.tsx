'use client';

import { Card } from '@/components/ui/card';
import { useConnections } from "@/hooks/use-connections";
import { ConnectionListSkeleton } from "@/components/ui/skeleton/connection-list";
import { useState } from 'react';
import { ConnectionListHeader } from './connection-list/header';
import { ConnectionCard } from './connection-list/connection-card';
import { BackupScheduleDialog } from './connection-list/backup-schedule-dialog';
import { useBackupConfigs } from './connection-list/use-backup-configs';
import type { SortBy } from './connection-list/types';

export function ConnectionsList() {
  const { connections, isLoading } = useConnections();
  const { backupConfigs, handleBackupNow, handleUpdateBackupConfig } = useBackupConfigs();
  const [searchQuery, setSearchQuery] = useState('');
  const [sortBy, setSortBy] = useState<SortBy>('name');
  const [scheduleDialogConnection, setScheduleDialogConnection] = useState<string | null>(null);

  const filteredConnections = connections?.filter(connection => 
    connection.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
    connection.host.toLowerCase().includes(searchQuery.toLowerCase())
  ).sort((a, b) => {
    switch (sortBy) {
      case 'name':
        return a.name.localeCompare(b.name);
      case 'status':
        return a.status.localeCompare(b.status);
      case 'type':
        return a.type.localeCompare(b.type);
      case 'lastBackup':
        const aDate = backupConfigs[a.id]?.lastBackup || '0';
        const bDate = backupConfigs[b.id]?.lastBackup || '0';
        return bDate.localeCompare(aDate);
      default:
        return 0;
    }
  });

  return (
    <Card className="col-span-3 backdrop-blur-xl bg-card/50">
      <div className="p-6">
        <ConnectionListHeader 
          searchQuery={searchQuery}
          onSearchChange={setSearchQuery}
          sortBy={sortBy}
          onSortChange={setSortBy}
        />
        
        {isLoading ? (
          <ConnectionListSkeleton />
        ) : (
          <div className="space-y-3">
            {filteredConnections?.map((connection) => (
              <ConnectionCard
                key={connection.id}
                connection={connection}
                backupConfig={backupConfigs[connection.id]}
                onBackupNow={handleBackupNow}
                onSchedule={() => setScheduleDialogConnection(connection.id)}
              />
            ))}
          </div>
        )}

        <BackupScheduleDialog
          connectionId={scheduleDialogConnection}
          connection={connections?.find(c => c.id === scheduleDialogConnection)}
          backupConfig={scheduleDialogConnection ? backupConfigs[scheduleDialogConnection] : undefined}
          onClose={() => setScheduleDialogConnection(null)}
          onUpdateConfig={handleUpdateBackupConfig}
        />
      </div>
    </Card>
  );
}