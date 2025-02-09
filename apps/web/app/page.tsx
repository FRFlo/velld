"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { StatsCards } from "@/components/views/dashboard/stats-card";
import { ActivityList } from "@/components/views/dashboard/activity-list";

export default function Home() {
  return (
    <div className="flex h-screen bg-background">
      <Sidebar />
      <div className="flex-1 overflow-x-hidden">
        <div className="h-full px-8 py-6 space-y-8">
          <div className="space-y-1">
            <h2 className="text-2xl font-semibold tracking-tight">Dashboard</h2>
            <p className="text-sm text-muted-foreground">
              Monitor and manage your database backups
            </p>
          </div>

          <StatsCards />

          <ActivityList />
        </div>
      </div>
    </div>
  );
}