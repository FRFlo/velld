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
    <Card className="max-w-2xl p-6">
      <div className="flex items-center justify-between mb-6">
        <div className="flex items-center space-x-4">
          <div className="p-2 rounded-full bg-muted/40">
            <User className="w-5 h-5 text-muted-foreground" />
          </div>
          {isLoadingProfile ? (
            <Skeleton className="h-6 w-32" />
          ) : (
            <h3 className="text-lg font-medium">{profile?.data.username}</h3>
          )}
        </div>
        <Button 
          variant="destructive" 
          size="sm"
          onClick={handleLogout}
          disabled={isLoading}
        >
          <LogOut className="w-4 h-4 mr-2" />
          {isLoading ? "Signing out..." : "Sign out"}
        </Button>
      </div>

      <div className="flex items-center justify-between p-4 bg-muted/40 rounded-lg">
        <div className="flex items-center space-x-4">
          <div className="p-2 rounded-full bg-background">
            <Shield className="w-4 h-4 text-muted-foreground" />
          </div>
          <div>
            <p className="font-medium">Password</p>
            <p className="text-sm text-muted-foreground">Change your account password</p>
          </div>
        </div>
        <Button variant="outline" size="sm">Change Password</Button>
      </div>
    </Card>
  );
}