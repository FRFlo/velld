"use client";

import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import { Bell, Mail, Webhook } from "lucide-react";
import { useSettings } from "@/hooks/use-settings";
import { Skeleton } from "@/components/ui/skeleton";
import { UpdateSettingsRequest } from "@/types/settings";

type SettingsValue = string | number | boolean;

export function NotificationSettings() {
  const { settings, isLoading, updateSettings, isUpdating } = useSettings();
  const [formValues, setFormValues] = useState<UpdateSettingsRequest>({});

  if (isLoading) {
    return (
      <Card className="max-w-2xl p-6">
        <div className="space-y-6">
          {/* Dashboard Notifications Skeleton */}
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <div className="p-2 rounded-full bg-muted/40">
                <Bell className="w-5 h-5 text-muted-foreground" />
              </div>
              <div className="space-y-2">
                <Skeleton className="h-5 w-40" />
                <Skeleton className="h-4 w-56" />
              </div>
            </div>
            <Skeleton className="h-6 w-12" />
          </div>

          {/* Email Notifications Skeleton */}
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="p-2 rounded-full bg-muted/40">
                  <Mail className="w-5 h-5 text-muted-foreground" />
                </div>
                <div className="space-y-2">
                  <Skeleton className="h-5 w-36" />
                  <Skeleton className="h-4 w-48" />
                </div>
              </div>
              <Skeleton className="h-6 w-12" />
            </div>
          </div>

          {/* Webhook Notifications Skeleton */}
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="p-2 rounded-full bg-muted/40">
                  <Webhook className="w-5 h-5 text-muted-foreground" />
                </div>
                <div className="space-y-2">
                  <Skeleton className="h-5 w-44" />
                  <Skeleton className="h-4 w-52" />
                </div>
              </div>
              <Skeleton className="h-6 w-12" />
            </div>
          </div>
        </div>
      </Card>
    );
  }

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (Object.keys(formValues).length > 0) {
      updateSettings(formValues);
    }
  };

  const handleSwitchChange = async (field: keyof UpdateSettingsRequest, checked: boolean) => {
    try {
      setFormValues(prev => ({ ...prev, [field]: checked }));
      await updateSettings({ [field]: checked });
    } catch (error) {
      // Revert the form value on error
      setFormValues(prev => ({ ...prev, [field]: !checked }));
      console.error(`Failed to update ${field}:`, error);
    }
  };

  const handleChange = (field: keyof UpdateSettingsRequest, value: SettingsValue) => {
    setFormValues(prev => ({ ...prev, [field]: value }));
  };

  return (
    <Card className="max-w-2xl p-6">
      <form onSubmit={handleSubmit}>
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <div className="p-2 rounded-full bg-muted/40">
                <Bell className="w-5 h-5 text-muted-foreground" />
              </div>
              <div>
                <h3 className="text-lg font-medium">Dashboard Notifications</h3>
                <p className="text-sm text-muted-foreground">
                  Receive notifications in the dashboard
                </p>
              </div>
            </div>
            <Switch
              checked={settings?.notify_dashboard ?? false}
              disabled={isUpdating}
              onCheckedChange={(checked) => handleSwitchChange('notify_dashboard', checked)}
            />
          </div>

          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="p-2 rounded-full bg-muted/40">
                  <Mail className="w-5 h-5 text-muted-foreground" />
                </div>
                <div>
                  <h3 className="text-lg font-medium">Email Notifications</h3>
                  <p className="text-sm text-muted-foreground">
                    Receive notifications via email
                  </p>
                </div>
              </div>
              <Switch
                checked={settings?.notify_email ?? false}
                disabled={isUpdating}
                onCheckedChange={(checked) => handleSwitchChange('notify_email', checked)}
              />
            </div>

            {settings?.notify_email && (
              <div className="ml-12 space-y-4">
                <div className="grid gap-2">
                  <Label>Email Address</Label>
                  <Input
                    type="email"
                    defaultValue={settings.email}
                    onChange={(e) => handleChange('email', e.target.value)}
                  />
                </div>
                <div className="grid gap-2">
                  <Label>SMTP Host</Label>
                  <Input
                    defaultValue={settings.smtp_host}
                    onChange={(e) => handleChange('smtp_host', e.target.value)}
                  />
                </div>
                <div className="grid gap-2">
                  <Label>SMTP Port</Label>
                  <Input
                    type="number"
                    defaultValue={settings.smtp_port}
                    onChange={(e) => handleChange('smtp_port', parseInt(e.target.value))}
                  />
                </div>
                <div className="grid gap-2">
                  <Label>SMTP Username</Label>
                  <Input
                    defaultValue={settings.smtp_username}
                    onChange={(e) => handleChange('smtp_username', e.target.value)}
                  />
                </div>
                <div className="grid gap-2">
                  <Label>SMTP Password</Label>
                  <Input
                    type="password"
                    placeholder="Enter new password to change"
                    onChange={(e) => handleChange('smtp_password', e.target.value)}
                  />
                </div>
              </div>
            )}
          </div>

          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div className="flex items-center space-x-4">
                <div className="p-2 rounded-full bg-muted/40">
                  <Webhook className="w-5 h-5 text-muted-foreground" />
                </div>
                <div>
                  <h3 className="text-lg font-medium">Webhook Notifications</h3>
                  <p className="text-sm text-muted-foreground">
                    Receive notifications via webhook
                  </p>
                </div>
              </div>
              <Switch
                checked={settings?.notify_webhook ?? false}
                disabled={isUpdating}
                onCheckedChange={(checked) => handleSwitchChange('notify_webhook', checked)}
              />
            </div>

            {settings?.notify_webhook && (
              <div className="ml-12 space-y-4">
                <div className="grid gap-2">
                  <Label>Webhook URL</Label>
                  <Input
                    defaultValue={settings.webhook_url}
                    onChange={(e) => handleChange('webhook_url', e.target.value)}
                  />
                </div>
              </div>
            )}
          </div>
        </div>

        <div className="mt-6 flex justify-end">
          <Button type="submit" disabled={isUpdating}>
            {isUpdating ? "Saving..." : "Save Changes"}
          </Button>
        </div>
      </form>
    </Card>
  );
}
