import { Card } from "@/components/ui/card";
import { Activity, Database, Server, Shield } from "lucide-react";
import { useConnections } from "@/hooks/use-connections";
import { formatSize } from "@/lib/helper";
import { Skeleton } from "@/components/ui/skeleton";

export function ConnectionStats() {
  const { stats, isLoading } = useConnections();

  const statsData = [
    {
      name: "Active Sources",
      value: stats?.total_connections ?? 0,
      icon: Database,
      description: "Connected databases",
      color: "text-blue-500",
    },
    {
      name: "Total Size",
      value: `${formatSize(stats?.average_size ?? 0)}`,
      icon: Server,
      description: "Across all databases",
      color: "text-emerald-500",
    },
    {
      name: "Secure Connections",
      value: `${stats?.ssl_percentage ?? 0}%`,
      icon: Shield,
      description: "SSL/TLS enabled",
      color: "text-purple-500",
    },
    {
      name: "Average Load",
      value: "0.5 ms",
      icon: Activity,
      description: "Response time",
      color: "text-amber-500",
    }
  ];

  return (
    <div className="grid gap-4 md:grid-cols-4">
      {statsData.map((stat) => {
        const Icon = stat.icon;
        return (
          <Card
            key={stat.name}
            className="p-6 bg-card/50 backdrop-blur-xl hover:bg-card/60 transition-colors"
          >
            <div className="flex items-center space-x-4">
              <div className={`p-2 rounded-md bg-primary/10 ${stat.color}`}>
                <Icon className="h-5 w-5" />
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground">
                  {stat.name}
                </p>
                {isLoading ? (
                  <Skeleton className="h-8 w-20 mt-1" />
                ) : (
                  <h3 className="text-2xl font-bold mt-1">{stat.value}</h3>
                )}
                <p className="text-xs text-muted-foreground mt-1">
                  {stat.description}
                </p>
              </div>
            </div>
          </Card>
        );
      })}
    </div>
  );
}