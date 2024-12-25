"use client";

import { LoginForm } from "@/components/views/auth/login-form";
import { AuthLayout } from "@/components/layout/auth";

export default function LoginPage() {
  return (
    <AuthLayout
      title="Welcome back"
      subtitle="Enter your credentials to access your account"
    >
      <LoginForm />
    </AuthLayout>
  );
}