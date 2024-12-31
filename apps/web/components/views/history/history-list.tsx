"use client";

import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Database, Download, FileDown, MoreVertical } from "lucide-react";
import { formatDistanceToNow } from "date-fns";

interface BackupHistory {
  id: string;
  database: string;
  type: "mysql" | "postgresql" | "mongodb";
  status: "completed" | "failed" | "running";
  size: string;
  duration: string;
  timestamp: Date;
  error?: string;
}

const history: BackupHistory[] = [
  {
    id: "1",
    database: "Production DB",
    type: "postgresql",
    status: "completed",
    size: "2.3 GB",
    duration: "4m 12s",
    timestamp: new Date(2024, 2, 20, 14, 30),
  },
  {
    id: "2",
    database: "Analytics",
    type: "mysql",
    status: "failed",
    size: "1.7 GB",
    duration: "2m 45s",
    timestamp: new Date(2024, 2, 20, 13, 15),
    error: "Connection timeout",
  },
  {
    id: "3",
    database: "User Data",
    type: "mongodb",
    status: "running",
    size: "4.1 GB",
    duration: "6m 20s",
    timestamp: new Date(2024, 2, 20, 12, 45),
  },
];

const statusColors = {
  completed: "bg-emerald-500/15 text-emerald-500",
  failed: "bg-red-500/15 text-red-500",
  running: "bg-blue-500/15 text-blue-500",
};

export function HistoryList() {
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
            {history.map((item) => (
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
                        <h4 className="font-medium">{item.database}</h4>
                        <Badge variant="secondary" className="text-xs">
                          {item.type}
                        </Badge>
                      </div>
                      <p className="text-sm text-muted-foreground mt-1">
                        {formatDistanceToNow(item.timestamp, { addSuffix: true })}
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center space-x-4">
                    <div className="text-right">
                      <Badge
                        variant="secondary"
                        className={statusColors[item.status]}
                      >
                        {item.status}
                      </Badge>
                      <p className="text-sm text-muted-foreground mt-1">
                        {item.duration}
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
                {item.error && (
                  <div className="mt-4 text-sm text-red-500">{item.error}</div>
                )}
              </div>
            ))}
          </div>
        </ScrollArea>
      </div>
    </Card>
  );
}