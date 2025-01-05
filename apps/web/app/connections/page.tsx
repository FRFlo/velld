"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { ConnectionsList } from "@/components/views/connections/connection-list";
import { ConnectionStats } from "@/components/views/connections/connection-stats";
import { AddConnectionDialog } from "@/components/views/connections/add-connection-dialog";

export default function ConnectionsPage() {
  return (
    <div className="flex h-screen bg-background/50 backdrop-blur-xl">
      <Sidebar />
      <div className="flex-1">
        <div className="h-full px-8 py-6">
          <div className="flex justify-between items-center mb-8">
            <div>
              <h2 className="text-3xl font-bold tracking-tight">Connections</h2>
              <p className="text-muted-foreground">
                Manage your database connections and sources
              </p>
            </div>
            <AddConnectionDialog/>
          </div>

          <div className="grid gap-6">
            <ConnectionStats />
            <ConnectionsList />
          </div>
        </div>
      </div>
    </div>
  );
}