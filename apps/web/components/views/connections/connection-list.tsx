'use client';

import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Database, ExternalLink, Power, Server, Settings } from 'lucide-react';
import { useConnections } from "@/hooks/use-connections";

const statusColors = {
  connected: 'bg-emerald-500/15 text-emerald-500',
  disconnected: 'bg-amber-500/15 text-amber-500',
  error: 'bg-red-500/15 text-red-500',
};

const typeLabels = {
  mysql: 'MySQL',
  postgresql: 'PostgreSQL',
  mongodb: 'MongoDB',
};

export function ConnectionsList() {
  const { connections, isLoading } = useConnections();

  return (
    <Card className="col-span-3 backdrop-blur-xl bg-card/50">
      <div className="p-6">
        <div className="flex justify-between items-center mb-6">
          <h3 className="text-lg font-semibold">Active Connections</h3>
          <Button variant="outline" size="sm">
            <Settings className="h-4 w-4 mr-2" />
            Manage All
          </Button>
        </div>
        <ScrollArea className="h-[500px] pr-4">
          {isLoading ? (
            <div>Loading...</div>
          ) : (
            <div className="space-y-4">
              {connections?.map((connection) => (
                <div
                  key={connection.id}
                  className="p-4 rounded-lg bg-background/50 hover:bg-background/60 transition-colors border"
                >
                  <div className="flex items-center justify-between">
                    <div className="flex items-center space-x-4">
                      <div className="p-2 rounded-md bg-primary/10">
                        <Database className="h-5 w-5 text-primary" />
                      </div>
                      <div>
                        <div className="flex items-center space-x-2">
                          <h4 className="font-medium">{connection.name}</h4>
                          <Badge variant="secondary" className="text-xs">
                            {typeLabels[connection.type]}
                          </Badge>
                        </div>
                        <p className="text-sm text-muted-foreground mt-1">
                          {connection.host}
                        </p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-4">
                      <div className="text-right">
                        <Badge
                          variant="secondary"
                          className={statusColors[connection.status]}
                        >
                          {connection.status}
                        </Badge>
                        <p className="text-sm text-muted-foreground mt-1">
                          {connection.lastSync}
                        </p>
                      </div>
                      <div className="flex space-x-2">
                        <Button variant="ghost" size="icon">
                          <ExternalLink className="h-4 w-4" />
                        </Button>
                        <Button variant="ghost" size="icon">
                          <Power className="h-4 w-4" />
                        </Button>
                      </div>
                    </div>
                  </div>
                  <div className="mt-4 flex items-center text-sm text-muted-foreground">
                    <div className="flex items-center">
                      <Server className="h-4 w-4 mr-1" />
                      Total Size: {connection.size}
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </ScrollArea>
      </div>
    </Card>
  );
}
