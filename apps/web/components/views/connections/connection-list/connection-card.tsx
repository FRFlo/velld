'use client';

import { formatSize } from '@/lib/helper';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Database, Clock, Settings2, Play, CheckCircle2 } from 'lucide-react';
import { cn } from '@/lib/utils';
import type { Connection } from '@/types/connection';
import { statusColors, typeLabels } from './types';
import { useBackup } from '@/hooks/use-backup';

interface ConnectionCardProps {
  connection: Connection;
  isScheduled?: boolean;
  lastBackupTime?: string;
  scheduleFrequency?: string;
  onSchedule: () => void;
}

export function ConnectionCard({
  connection,
  isScheduled,
  lastBackupTime,
  scheduleFrequency,
  onSchedule,
}: ConnectionCardProps) {
  const { createBackup, isCreating } = useBackup();

  return (
    <div
      className={cn(
        "p-4 rounded-lg transition-all duration-200 border",
        isScheduled
          ? "bg-emerald-500/5 hover:bg-emerald-500/10 border-emerald-500/20"
          : "bg-background/50 hover:bg-background/80 border-border/50"
      )}
    >
      <div className="flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <div className={cn(
            "p-2 rounded-md",
            isScheduled
              ? "bg-emerald-500/10"
              : "bg-primary/10"
          )}>
            <Database className={cn(
              "h-5 w-5",
              isScheduled
                ? "text-emerald-500"
                : "text-primary"
            )} />
          </div>
          <div>
            <div className="flex items-center space-x-2">
              <h4 className="font-medium">{connection.name}</h4>
              <Badge variant="outline" className="text-xs">
                {typeLabels[connection.type]}
              </Badge>
              <Badge
                variant="outline"
                className={cn("border", statusColors[connection.status])}
              >
                {connection.status}
              </Badge>
            </div>
            <div className="flex items-center space-x-4 text-sm text-muted-foreground mt-1.5">
              <span>{connection.host}</span>
              <span>•</span>
              <span>{formatSize(connection.database_size)}</span>
              {isScheduled && scheduleFrequency && (
                <>
                  <span>•</span>
                  <span className="flex items-center text-emerald-500">
                    <CheckCircle2 className="h-3 w-3 mr-1" />
                    Auto-backup {scheduleFrequency}
                  </span>
                </>
              )}
              {lastBackupTime && (
                <>
                  <span>•</span>
                  <span className="flex items-center">
                    <Clock className="h-3 w-3 mr-1" />
                    Last backup: {new Date(lastBackupTime).toLocaleString()}
                  </span>
                </>
              )}
            </div>
          </div>
        </div>
        
        <div className="flex items-center space-x-2">
          <Button 
            variant={isScheduled ? "default" : "secondary"}
            size="sm"
            onClick={() => createBackup(connection.id)}
            disabled={isCreating}
            className="space-x-1"
          >
            <Play className="h-3 w-3" />
            <span>Backup Now</span>
          </Button>
          
          <Button 
            variant={isScheduled ? "default" : "outline"} 
            size="sm" 
            className="space-x-1"
            onClick={onSchedule}
          >
            <Settings2 className="h-3 w-3" />
            <span>Schedule</span>
          </Button>
        </div>
      </div>
    </div>
  );
} 