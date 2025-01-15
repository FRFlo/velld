"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { HistoryStats } from "@/components/views/history/history-stats";
import { HistoryList } from "@/components/views/history/history-list";

export default function HistoryPage() {
  return (
    <div className="flex h-screen bg-background/50 backdrop-blur-xl">
      <Sidebar />
      <div className="flex-1">
        <div className="h-full px-8 py-6">
          <div className="mb-8">
            <h2 className="text-3xl font-bold tracking-tight">Backup History</h2>
            <p className="text-muted-foreground">
              View and manage your backup history across all databases
            </p>
          </div>

          <div className="grid gap-6">
            <HistoryStats />
            <HistoryList />
          </div>
        </div>
      </div>
    </div>
  );
}