"use client";

import { Layout, Database, Server, Settings, History, Shield } from "lucide-react";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { usePathname } from "next/navigation";
import Link from "next/link";

const navigation = [
  { name: "Dashboard", href: "/", icon: Layout },
  { name: "Connections", href: "/connections", icon: Server },
  { name: "Backup History", href: "/history", icon: History },
  { name: "Security", href: "/security", icon: Shield },
  { name: "Settings", href: "/settings", icon: Settings },
];

export function Sidebar() {
  const pathname = usePathname();

  return (
    <div className="w-64 border-r bg-card/50 backdrop-blur-xl p-6 space-y-6">
      <div className="flex items-center space-x-2">
        <Database className="h-6 w-6 text-primary" />
        <h1 className="text-xl font-semibold">Velld</h1>
      </div>
      <nav className="space-y-1">
        {navigation.map((item) => {
          const Icon = item.icon;
          return (
            <Link key={item.name} href={item.href}>
              <Button
                variant={pathname === item.href ? "secondary" : "ghost"}
                className={cn(
                  "w-full justify-start",
                  pathname === item.href && "bg-secondary"
                )}
              >
                <Icon className="mr-2 h-4 w-4" />
                {item.name}
              </Button>
            </Link>
          );
        })}
      </nav>
    </div>
  );
}