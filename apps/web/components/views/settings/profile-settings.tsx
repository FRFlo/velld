"use client";

import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { LogOut, Shield, User } from "lucide-react";
import { useProfile } from "@/hooks/use-profile";
import { logout } from "@/lib/helper";
import { Skeleton } from "@/components/ui/skeleton";

export function ProfileSettings() {
  const [isLoading, setIsLoading] = useState(false);
  const { data: profile, isLoading: isLoadingProfile } = useProfile();

  const handleLogout = () => {
    setIsLoading(true);
    logout();
  };

  return (
    <div className="space-y-6 max-w-2xl">
      <Card className="p-6">
        <h3 className="text-lg font-medium mb-4">Profile Information</h3>
        <div className="space-y-4">
          <div className="flex items-center space-x-4 p-4 bg-muted/40 rounded-lg">
            <div className="p-2 rounded-full bg-background">
              <User className="w-6 h-6 text-muted-foreground" />
            </div>
            <div>
              {isLoadingProfile ? (
                <Skeleton className="h-8 w-32" />
              ) : (
                <p className="font-medium">{profile?.data.username}</p>
              )}
            </div>
          </div>
          {/* <div className="space-y-4">
            <div className="space-y-2">
              <Label>Username</Label>
              <Input 
                value={profile?.username || ''} 
                readOnly 
                disabled={isLoadingProfile}
              />
            </div>
          </div> */}
        </div>
      </Card>

      <Card className="p-6">
        <h3 className="text-lg font-medium mb-4">Security</h3>
        <div className="space-y-4">
          <div className="flex-none sm:flex items-center justify-between p-4 bg-muted/40 rounded-lg">
            <div className="flex items-center space-x-4 mb-4 sm:mb-0">
              <div className="p-2 rounded-full bg-background">
                <Shield className="w-5 h-5 text-muted-foreground" />
              </div>
              <div>
                <p className="font-medium">Password</p>
                <p className="text-sm text-muted-foreground">Last changed 2 months ago</p>
              </div>
            </div>
            <Button variant="outline" size="sm">Change Password</Button>
          </div>
        </div>
      </Card>

      <Card className="p-6 border-destructive/20">
        <h3 className="text-lg font-medium mb-4">Account Actions</h3>
        <div className="space-y-4">
          <Button 
            variant="destructive" 
            className="w-full sm:w-auto" 
            onClick={handleLogout}
            disabled={isLoading}
          >
            <LogOut className="w-4 h-4 mr-2" />
            {isLoading ? "Signing out..." : "Sign out"}
          </Button>
        </div>
      </Card>
    </div>
  );
}