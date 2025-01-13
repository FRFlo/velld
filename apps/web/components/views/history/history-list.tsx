"use client";

import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Database, Download, FileDown, MoreVertical } from "lucide-react";
import { formatDistanceToNow, parseISO } from "date-fns";
import { useBackup } from "@/hooks/use-backup";
import { BackupList } from "@/types/backup";
import { HistoryListSkeleton } from "@/components/ui/skeleton/history-list";
import { calculateDuration, formatSize } from "@/lib/helper";

const statusColors = {
  completed: "bg-emerald-500/15 text-emerald-500",
  failed: "bg-red-500/15 text-red-500",
  running: "bg-blue-500/15 text-blue-500",
};

export function HistoryList() {
  const { backups, isLoading } = useBackup();

  if (isLoading) {
    return <HistoryListSkeleton />;
  }

  return (
    <Card className="col-span-3 backdrop-blur-xl bg-card/50">
      <div className="p-6">
        <div className="flex justify-between items-center mb-6">
          <h3 className="text-lg font-semibold">Recent Backups</h3>
          <Button variant="outline" size="sm">
            <FileDown className="h-4 w-4 mr-2" />
            Export Log
          </Button>
        </div>
        <ScrollArea className="h-[400px] pr-4">
          <div className="space-y-4">
            {backups?.map((item: BackupList) => (
              <div
                key={item.id}
                className="p-4 rounded-lg bg-background/50 hover:bg-background/60 transition-colors border"
              >
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-4">
                    <div className="p-2 rounded-md bg-primary/10">
                      <Database className="h-5 w-5 text-primary" />
                    </div>
                    <div>
                      <div className="flex items-center space-x-2">
                        <h4 className="font-medium">{item.path.split('\\').pop()}</h4>
                        <Badge variant="secondary" className="text-xs">
                          {item.database_type}
                        </Badge>
                      </div>
                      <p className="text-sm text-muted-foreground mt-1">
                        {formatDistanceToNow(parseISO(item.created_at), { addSuffix: true })} | {formatSize(item.size)}
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-4">
                    <div className="text-right">
                      <Badge
                        variant="secondary"
                        className={statusColors[item.status as keyof typeof statusColors]}
                      >
                        {item.status}
                      </Badge>
                      <p className="text-sm text-muted-foreground mt-1">
                        {calculateDuration(item.started_time, item.completed_time)}
                      </p>
                    </div>
                    <div className="flex space-x-2">
                      <Button variant="ghost" size="icon">
                        <Download className="h-4 w-4" />
                      </Button>
                      <Button variant="ghost" size="icon">
                        <MoreVertical className="h-4 w-4" />
                      </Button>
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </ScrollArea>
      </div>
    </Card>
  );
}