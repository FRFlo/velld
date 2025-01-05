"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { StatsCards } from "@/components/views/dashboard/stats-card";
import { BackupList } from "@/components/views/backup/backup-list";
import { Card } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "@/components/ui/button";
import { AddBackupDialog } from "@/components/views/backup/add-backup-dialog";

export default function Home() {
  return (
    <div className="flex h-screen bg-background">
      <Sidebar />
      <div className="flex-1 overflow-hidden">
        <div className="h-full px-8 py-6 space-y-8">
          <div className="flex justify-between items-center">
            <div className="space-y-1">
              <h2 className="text-2xl font-semibold tracking-tight">Dashboard</h2>
              <p className="text-sm text-muted-foreground">
                Monitor and manage your database backups
              </p>
            </div>
            <AddBackupDialog />
          </div>

          <div className="space-y-6">
            <StatsCards />
            <Card className="p-0 border-none shadow-none bg-transparent">
              <Tabs defaultValue="recent" className="w-full">
                <div className="flex justify-between items-center mb-4">
                  <TabsList className="h-8">
                    <TabsTrigger value="recent" className="text-xs">Recent Activity</TabsTrigger>
                    <TabsTrigger value="scheduled" className="text-xs">Scheduled</TabsTrigger>
                  </TabsList>
                  <Button variant="ghost" size="sm" className="h-8 text-xs">
                    View All
                  </Button>
                </div>
                <TabsContent value="recent" className="m-0">
                  <BackupList />
                </TabsContent>
                <TabsContent value="scheduled" className="m-0">
                  <BackupList />
                </TabsContent>
              </Tabs>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
}