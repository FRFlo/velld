"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Link from "next/link";

export function LoginForm() {
  const [isLoading, setIsLoading] = useState(false);

  async function onSubmit(event: React.FormEvent) {
    event.preventDefault();
    setIsLoading(true);
    
    setTimeout(() => {
      setIsLoading(false);
    }, 1000);
  }

  return (
    <form onSubmit={onSubmit} className="space-y-6">
      <div className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="username">Username</Label>
          <Input
            id="username"
            placeholder="Enter your username"
            type="text"
            disabled={isLoading}
            required
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            placeholder="Enter your password"
            type="password"
            disabled={isLoading}
            required
          />
        </div>
      </div>
      
      <Button className="w-full" type="submit" disabled={isLoading}>
        {isLoading ? "Signing in..." : "Sign in"}
      </Button>

      <div className="text-center text-sm">
        <span className="text-muted-foreground">Dont have an account? </span>
        <Link 
          href="/register" 
          className="font-medium hover:text-primary underline-offset-4 hover:underline"
        >
          Sign up
        </Link>
      </div>
    </form>
  );
}