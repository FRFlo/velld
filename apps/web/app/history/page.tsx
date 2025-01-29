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
          <div className="flex flex-col gap-6">
            <HistoryStats />
            <HistoryList />
          </div>
        </div>
      </div>
    </div>
  );
}