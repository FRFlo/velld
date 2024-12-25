"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/hooks/use-toast";
import { useRegister } from "@/hooks/use-auth";
import Link from "next/link";

export function RegisterForm() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    confirmPassword: "",
  });
  
  const { toast } = useToast();
  const { mutate: register, isPending } = useRegister();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.id]: e.target.value,
    });
  };

  async function onSubmit(event: React.FormEvent) {
    event.preventDefault();
    
    if (!formData.username || !formData.password || !formData.confirmPassword) {
      toast({
        title: "Error",
        description: "All fields are required",
        variant: "destructive",
      });
      return;
    }

    if (formData.password.length < 6) {
      toast({
        title: "Error",
        description: "Password must be at least 6 characters long",
        variant: "destructive",
      });
      return;
    }

    if (formData.password !== formData.confirmPassword) {
      toast({
        title: "Error",
        description: "Passwords do not match",
        variant: "destructive",
      });
      return;
    }

    register({
      username: formData.username,
      password: formData.password,
    });
  }

  return (
    <form onSubmit={onSubmit} className="space-y-6">
      <div className="space-y-4">
        <div className="space-y-2">
          <Label htmlFor="username">Username</Label>
          <Input
            id="username"
            placeholder="Choose a username"
            type="text"
            disabled={isPending}
            required
            value={formData.username}
            onChange={handleChange}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            placeholder="Choose a password"
            type="password"
            disabled={isPending}
            required
            value={formData.password}
            onChange={handleChange}
          />
        </div>
        <div className="space-y-2">
          <Label htmlFor="confirmPassword">Confirm Password</Label>
          <Input
            id="confirmPassword"
            placeholder="Confirm your password"
            type="password"
            disabled={isPending}
            required
            value={formData.confirmPassword}
            onChange={handleChange}
          />
        </div>
      </div>
      
      <Button className="w-full" type="submit" disabled={isPending}>
        {isPending ? "Signing up..." : "Sign up"}
      </Button>

      <div className="text-center text-sm">
        <span className="text-muted-foreground">Already have an account? </span>
        <Link 
          href="/login" 
          className="font-medium hover:text-primary underline-offset-4 hover:underline"
        >
          Sign in
        </Link>
      </div>
    </form>
  );
}