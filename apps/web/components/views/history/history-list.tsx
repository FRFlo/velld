"use client";

import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Database, Download, FileDown, MoreVertical } from "lucide-react";
import { formatDistanceToNow, parseISO } from "date-fns";
import { useBackup } from "@/hooks/use-backup";
import { BackupList } from "@/types/backup";
import { HistoryListSkeleton } from "@/components/ui/skeleton/history-list";
import { calculateDuration, formatSize } from "@/lib/helper";
import { HistoryFilters } from "./history-filters";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { ScrollArea } from "@/components/ui/scroll-area";
import { Bell } from "lucide-react";

// Sample notification data
const notifications = [
  {
    id: 1,
    title: "Backup Completed",
    message: "Database backup for 'Production DB' completed successfully",
    time: "2 minutes ago",
    status: "completed"
  },
  {
    id: 2,
    title: "Backup Failed",
    message: "Backup failed for 'Test DB' due to connection error",
    time: "1 hour ago",
    status: "failed"
  },
  {
    id: 3,
    title: "Backup Started",
    message: "New backup started for 'Development DB'",
    time: "2 hours ago",
    status: "running"
  },
];

const statusColors = {
  completed: "bg-emerald-500/15 text-emerald-500",
  failed: "bg-red-500/15 text-red-500",
  running: "bg-blue-500/15 text-blue-500",
};

export function HistoryList() {
  const { backups, isLoading, pagination, page, setPage } = useBackup();
  const totalPages = pagination?.total_pages || 1;

  return (
    <Card className="col-span-3 backdrop-blur-xl bg-card/50">
      <div className="flex flex-col h-full">
        <div className="p-6 pb-0">
          <div className="flex justify-between items-center mb-6">
            <h3 className="text-lg font-semibold">Recent Backups</h3>
            <Button variant="outline" size="sm">
              <FileDown className="h-4 w-4 mr-2" />
              Export Log
            </Button>
          </div>
          <HistoryFilters />
        </div>
        <div className="flex flex-row p-6 pt-4 space-x-4">
          <div className="flex-1 space-y-4">
            {isLoading ? (
              <HistoryListSkeleton />
            ) : (
              backups?.map((item: BackupList) => (
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
              ))
            )}
          </div>

          {/* Notification sidebar */}
          <div className="w-80 border-l pl-4">
            <div className="flex items-center justify-between mb-4">
              <h4 className="text-sm font-semibold">Notifications</h4>
              <Badge variant="secondary" className="text-xs">
                {notifications.length} new
              </Badge>
            </div>
            <ScrollArea className="h-[600px]">
              <div className="space-y-4 pr-4">
                {notifications.map((notification) => (
                  <div
                    key={notification.id}
                    className="p-3 rounded-lg bg-background/50 border"
                  >
                    <div className="flex items-start space-x-3">
                      <div className="p-1.5 rounded-md bg-primary/10">
                        <Bell className="h-4 w-4 text-primary" />
                      </div>
                      <div>
                        <h5 className="text-sm font-medium">{notification.title}</h5>
                        <p className="text-xs text-muted-foreground mt-1">
                          {notification.message}
                        </p>
                        <p className="text-xs text-muted-foreground mt-2">
                          {notification.time}
                        </p>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            </ScrollArea>
          </div>
        </div>

        {pagination && (
          <div className="p-6">
            <Pagination>
              <PaginationContent>
                <PaginationItem>
                  <PaginationPrevious 
                    href="#"
                    onClick={(e) => {
                      e.preventDefault();
                      if (page > 1) setPage(page - 1);
                    }}
                    className={page <= 1 ? 'pointer-events-none opacity-50' : ''}
                  />
                </PaginationItem>
                
                {Array.from({ length: totalPages }, (_, i) => i + 1).map((pageNum) => (
                  <PaginationItem key={pageNum}>
                    <PaginationLink
                      href="#"
                      isActive={pageNum === page}
                      onClick={(e) => {
                        e.preventDefault();
                        setPage(pageNum);
                      }}
                    >
                      {pageNum}
                    </PaginationLink>
                  </PaginationItem>
                ))}

                <PaginationItem>
                  <PaginationNext 
                    href="#"
                    onClick={(e) => {
                      e.preventDefault();
                      if (page < totalPages) setPage(page + 1);
                    }}
                    className={page >= totalPages ? 'pointer-events-none opacity-50' : ''}
                  />
                </PaginationItem>
              </PaginationContent>
            </Pagination>
          </div>
        )}
      </div>
    </Card>
  );
}