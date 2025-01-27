import { Skeleton } from "@/components/ui/skeleton";
import { Card } from "@/components/ui/card";

export function BackupStatsSkeleton() {
  return (
    <div className="grid gap-4 md:grid-cols-4">
      {[1, 2, 3, 4].map((i) => (
        <Card
          key={i}
          className="p-6 bg-background border border-border/50"
        >
          <div className="flex justify-between items-start">
            <div className="space-y-2">
              <Skeleton className="h-4 w-24" />
              <Skeleton className="h-7 w-16" />
            </div>
            <Skeleton className="h-4 w-10" />
          </div>
          <div className="mt-4">
            <Skeleton className="h-1.5 w-full" />
          </div>
        </Card>
      ))}
    </div>
  );
} 