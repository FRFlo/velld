import { Card } from "@/components/ui/card";
import { ArrowUpIcon, ArrowDownIcon } from "lucide-react";

const stats = [
  {
    name: "Success Rate",
    value: "99.8%",
    change: "+0.3%",
    status: "Highly reliable performance",
    trend: "up",
  },
  {
    name: "Failed Backups",
    value: "2",
    change: "+1",
    status: "Requires attention",
    trend: "down",
  },
  {
    name: "Average Duration",
    value: "4.2m",
    change: "75%",
    status: "Processing normally",
    showProgress: true,
  },
  {
    name: "Storage Used",
    value: "2.1TB",
    change: "82%",
    status: "Sufficient capacity",
    showProgress: true,
  },
];

export function HistoryStats() {
  return (
    <div className="grid gap-4 md:grid-cols-4">
      {stats.map((stat) => (
        <Card
          key={stat.name}
          className="p-6 bg-background border border-border/50 hover:border-border/80 transition-colors"
        >
          <div className="flex justify-between items-start">
            <div className="space-y-2">
              <p className="text-sm font-medium text-muted-foreground">
                {stat.name}
              </p>
              <p className="text-2xl font-semibold">{stat.value}</p>
            </div>
            <div className={`flex items-center ${
              stat.trend === 'up' ? 'text-emerald-500' : 
              stat.trend === 'down' ? 'text-amber-500' : 
              'text-blue-500'
            } text-xs font-medium`}>
              {stat.trend && (
                <>
                  {stat.trend === 'up' ? (
                    <ArrowUpIcon className="h-3 w-3 mr-1" />
                  ) : (
                    <ArrowDownIcon className="h-3 w-3 mr-1" />
                  )}
                </>
              )}
              {stat.change}
            </div>
          </div>
          {stat.showProgress ? (
            <div className="mt-4">
              <div className="h-1 bg-muted/30 rounded-full overflow-hidden">
                <div 
                  className="h-full bg-blue-500/50 rounded-full" 
                  style={{ width: stat.change }}
                />
              </div>
            </div>
          ) : (
            <div className="mt-4 flex items-center text-[13px] text-muted-foreground/80">
              <div className={`h-1.5 w-1.5 rounded-full ${
                stat.trend === 'up' ? 'bg-emerald-500' : 
                stat.trend === 'down' ? 'bg-amber-500' : 
                'bg-blue-500'
              } mr-2`} />
              {stat.status}
            </div>
          )}
        </Card>
      ))}
    </div>
  );
}

export default HistoryStats;