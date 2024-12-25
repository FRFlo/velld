"use client";

import { Sidebar } from "@/components/layout/sidebar";
import { SettingsLayout } from "@/components/layout/setting";
import { ProfileSettings } from "@/components/views/settings/profile-settings";

export default function SettingsPage() {
  return (
    <div className="flex h-screen bg-background/50 backdrop-blur-xl">
      <Sidebar />
      <SettingsLayout>
        <ProfileSettings />
      </SettingsLayout>
    </div>
  );
}