import { Database } from "lucide-react";
import { ReactNode } from "react";

interface AuthLayoutProps {
  children: ReactNode;
  title: string;
  subtitle: string;
}

export function AuthLayout({ children, title, subtitle }: AuthLayoutProps) {
  return (
    <div className="min-h-screen grid lg:grid-cols-2">
      {/* Left: Auth Form */}
      <div className="flex items-center justify-center p-8">
        <div className="w-full max-w-sm space-y-8">
          <div className="space-y-6">
            <div className="flex items-center space-x-2">
              <Database className="h-6 w-6" />
              <span className="text-xl font-semibold">Velld</span>
            </div>
            <div className="space-y-2">
              <h1 className="text-2xl font-semibold tracking-tight">{title}</h1>
              <p className="text-sm text-muted-foreground">{subtitle}</p>
            </div>
          </div>
          {children}
        </div>
      </div>
      
      {/* Right: Background Pattern */}
      <div className="hidden lg:block bg-muted/40">
        <div className="absolute inset-0 bg-grid-black/[0.02] bg-[size:24px_24px]" />
      </div>
    </div>
  );
}