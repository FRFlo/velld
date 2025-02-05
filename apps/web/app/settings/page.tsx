"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { ProfileSettings } from "@/components/views/settings/profile-settings";
import { NotificationSettings } from "@/components/views/settings/notification-settings";


export default function SettingsPage() {
  return (
    <div className="flex-none sm:flex h-screen bg-background/50 backdrop-blur-xl">
      <Sidebar />
      <div className="flex-1">
        <div className="h-full px-8 py-6 space-y-6">
          <ProfileSettings />
          <NotificationSettings />
        </div>
      </div>
    </div>
  );
}