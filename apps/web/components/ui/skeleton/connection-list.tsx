import { Skeleton } from "@/components/ui/skeleton";
import { Card } from "@/components/ui/card";

export function ConnectionListSkeleton() {
  return (
    <Card>
      <div className="p-6">
        <div className="space-y-4">
          <div className="flex items-center justify-between">
            <Skeleton className="h-8 w-[200px]" />
            <Skeleton className="h-8 w-[100px]" />
          </div>
          
          {[1, 2, 3].map((i) => (
            <div key={i} className="border rounded-lg p-4">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-4">
                  <Skeleton className="h-12 w-12 rounded-full" />
                  <div className="space-y-2">
                    <Skeleton className="h-4 w-[200px]" />
                    <Skeleton className="h-4 w-[150px]" />
                  </div>
                </div>
                <div className="flex items-center space-x-4">
                  <div className="text-right space-y-2">
                    <Skeleton className="h-5 w-[100px]" />
                    <Skeleton className="h-4 w-[80px]" />
                  </div>
                  <div className="flex space-x-2">
                    <Skeleton className="h-8 w-8" />
                    <Skeleton className="h-8 w-8" />
                  </div>
                </div>
              </div>
              <div className="mt-4">
                <Skeleton className="h-4 w-[120px]" />
              </div>
            </div>
          ))}
        </div>
      </div>
    </Card>
  );
}