import { Card } from "@/components/ui/card";
import { ArrowUpIcon, ArrowDownIcon } from "lucide-react";

export function StatsCards() {
  return (
    <div className="grid gap-4 md:grid-cols-3">
      <Card className="p-6 bg-background border border-border/50 hover:border-border/80 transition-colors">
        <div className="flex justify-between items-start">
          <div className="space-y-2">
            <p className="text-sm font-medium text-muted-foreground">
              Total Backups
            </p>
            <p className="text-2xl font-semibold">128</p>
          </div>
          <div className="flex items-center text-emerald-500 text-xs font-medium">
            <ArrowUpIcon className="h-3 w-3 mr-1" />
            12%
          </div>
        </div>
        <div className="mt-4 flex items-center text-[13px] text-muted-foreground/80">
          <div className="h-1.5 w-1.5 rounded-full bg-emerald-500 mr-2" />
          All systems operational
        </div>
      </Card>
      <Card className="p-6 bg-background border border-border/50 hover:border-border/80 transition-colors">
        <div className="flex justify-between items-start">
          <div className="space-y-2">
            <p className="text-sm font-medium text-muted-foreground">
              Active Connections
            </p>
            <p className="text-2xl font-semibold">7</p>
          </div>
          <div className="flex items-center text-amber-500 text-xs font-medium">
            <ArrowDownIcon className="h-3 w-3 mr-1" />
            2
          </div>
        </div>
        <div className="mt-4 flex items-center text-[13px] text-muted-foreground/80">
          <div className="h-1.5 w-1.5 rounded-full bg-amber-500 mr-2" />
          3 pending syncs
        </div>
      </Card>
      <Card className="p-6 bg-background border border-border/50 hover:border-border/80 transition-colors">
        <div className="flex justify-between items-start">
          <div className="space-y-2">
            <p className="text-sm font-medium text-muted-foreground">
              Storage Used
            </p>
            <p className="text-2xl font-semibold">256GB</p>
          </div>
          <div className="text-xs font-medium text-blue-500">75%</div>
        </div>
        <div className="mt-4">
          <div className="h-1 bg-muted/30 rounded-full overflow-hidden">
            <div className="h-full bg-blue-500/50 rounded-full w-3/4" />
          </div>
        </div>
      </Card>
    </div>
  );
}