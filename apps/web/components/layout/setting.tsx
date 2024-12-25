import { ReactNode } from "react";

interface SettingsLayoutProps {
  children: ReactNode;
}

export function SettingsLayout({ children }: SettingsLayoutProps) {
  return (
    <div className="flex-1 overflow-hidden relative">
      <div className="h-full px-8 py-6">
        <div className="mb-8">
          <h2 className="text-3xl font-bold tracking-tight">Settings</h2>
          <p className="text-muted-foreground">
            Manage your account settings and preferences
          </p>
        </div>
        {children}
      </div>
    </div>
  );
}