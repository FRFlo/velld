"use client";

import { RegisterForm } from "@/components/views/auth/register-form";
import { AuthLayout } from "@/components/layout/auth";

export default function RegisterPage() {
  return (
    <AuthLayout
      title="Create an account"
      subtitle="Enter your details to get started"
    >
      <RegisterForm />
    </AuthLayout>
  );
}