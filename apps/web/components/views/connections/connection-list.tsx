'use client';

import { Card } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { ScrollArea } from '@/components/ui/scroll-area';
import { Database, Settings2, ExternalLink, Power, Server } from 'lucide-react';

interface Connection {
  id: string;
  name: string;
  type: 'mysql' | 'postgresql' | 'mongodb';
  host: string;
  status: 'connected' | 'disconnected' | 'error';
  lastSync: string;
  size: string;
}

const connections: Connection[] = [
  {
    id: '1',
    name: 'Production Database',
    type: 'postgresql',
    host: 'db-prod-01.example.com',
    status: 'connected',
    lastSync: '2 minutes ago',
    size: '789 GB',
  },
  {
    id: '2',
    name: 'Analytics DB',
    type: 'mysql',
    host: 'analytics-db.example.com',
    status: 'connected',
    lastSync: '5 minutes ago',
    size: '234 GB',
  },
  {
    id: '3',
    name: 'User Data Store',
    type: 'mongodb',
    host: 'mongo-cluster-1.example.com',
    status: 'error',
    lastSync: 'Failed 10 min ago',
    size: '1.2 TB',
  },
];

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
  return (
    <Card className="col-span-3 backdrop-blur-xl bg-card/50">
      <div className="p-6">
        <div className="flex justify-between items-center mb-6">
          <h3 className="text-lg font-semibold">Active Connections</h3>
          <Button variant="outline" size="sm">
            <Settings2 className="h-4 w-4 mr-2" />
            Manage All
          </Button>
        </div>
        <ScrollArea className="h-[500px] pr-4">
          <div className="space-y-4">
            {connections.map((connection) => (
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
        </ScrollArea>
      </div>
    </Card>
  );
}
