"use client";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Clock, Database, HardDrive, MoreHorizontal } from "lucide-react";

interface BackupItem {
  id: number;
  name: string;
  source: string;
  size: string;
  status: "completed" | "pending" | "failed";
  time: string;
  type: "mysql" | "postgresql" | "mongodb";
}

const backups: BackupItem[] = [
  {
    id: 1,
    name: "Production DB",
    source: "MySQL Production",
    size: "2.3GB",
    status: "completed",
    time: "2 hours ago",
    type: "mysql",
  },
  {
    id: 2,
    name: "Analytics DB",
    source: "PostgreSQL Analytics",
    size: "1.7GB",
    status: "pending",
    time: "Scheduled in 1 hour",
    type: "postgresql",
  },
  {
    id: 3,
    name: "User Data",
    source: "MongoDB Users",
    size: "4.1GB",
    status: "completed",
    time: "5 hours ago",
    type: "mongodb",
  },
];

const statusColors = {
  completed: "bg-emerald-500/15 text-emerald-500 border-emerald-500/20",
  pending: "bg-amber-500/15 text-amber-500 border-amber-500/20",
  failed: "bg-red-500/15 text-red-500 border-red-500/20",
};

export function BackupList() {
  return (
    <ScrollArea className="h-[400px]">
      <div className="space-y-2">
        {backups.map((backup) => (
          <div
            key={backup.id}
            className="group p-3 rounded-lg border border-border/50 hover:border-border/80 transition-colors"
          >
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="p-2 rounded-md bg-primary/5">
                  <Database className="h-4 w-4 text-primary" />
                </div>
                <div>
                  <div className="flex items-center space-x-2">
                    <p className="text-sm font-medium">{backup.name}</p>
                    <Badge variant="outline" className="text-[10px] font-normal px-1 h-4">
                      {backup.type}
                    </Badge>
                  </div>
                  <div className="flex items-center space-x-2 text-xs text-muted-foreground mt-0.5">
                    <HardDrive className="h-3 w-3" />
                    <span>{backup.source}</span>
                    <span className="text-muted-foreground/40">•</span>
                    <span>{backup.size}</span>
                  </div>
                </div>
              </div>
              <div className="flex items-center space-x-4">
                <Badge 
                  variant="outline" 
                  className={`text-[10px] ${statusColors[backup.status]}`}
                >
                  {backup.status}
                </Badge>
                <div className="flex items-center text-xs text-muted-foreground">
                  <Clock className="h-3 w-3 mr-1" />
                  {backup.time}
                </div>
                <Button 
                  variant="ghost" 
                  size="icon" 
                  className="h-6 w-6 opacity-0 group-hover:opacity-100 transition-opacity"
                >
                  <MoreHorizontal className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </ScrollArea>
  );
}