import { Card } from "@/components/ui/card";
import { CheckCircle2, AlertCircle, Clock, HardDrive } from "lucide-react";

const stats = [
  {
    name: "Success Rate",
    value: "99.8%",
    icon: CheckCircle2,
    description: "Last 30 days",
    color: "text-emerald-500",
  },
  {
    name: "Failed Backups",
    value: "2",
    icon: AlertCircle,
    description: "Requires attention",
    color: "text-red-500",
  },
  {
    name: "Average Duration",
    value: "4.2m",
    icon: Clock,
    description: "Per backup",
    color: "text-blue-500",
  },
  {
    name: "Storage Used",
    value: "2.1TB",
    icon: HardDrive,
    description: "Total backups",
    color: "text-purple-500",
  },
];

export function HistoryStats() {
  return (
    <div className="grid gap-4 md:grid-cols-4">
      {stats.map((stat) => {
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
                <h3 className="text-2xl font-bold mt-1">{stat.value}</h3>
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