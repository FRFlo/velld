'use client';

import { useState } from 'react';
import { useConnections } from "@/hooks/use-connections";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Button } from '@/components/ui/button';
import { Label } from '@/components/ui/label';
import { Input } from '@/components/ui/input';
import { Switch } from '@/components/ui/switch';

interface BackupFormProps {
  onSuccess?: () => void;
  onCancel?: () => void;
}

interface BackupFormData {
  connectionId: string;
  name: string;
  description?: string;
  scheduled: boolean;
  cronExpression?: string;
  compression: boolean;
}

export function BackupForm({ onSuccess, onCancel }: BackupFormProps) {
  const [formData, setFormData] = useState<BackupFormData>({
    connectionId: '',
    name: '',
    description: '',
    scheduled: false,
    cronExpression: '',
    compression: true,
  });

  const { connections } = useConnections();
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);
    // TODO: Implement backup creation logic
    try {
      // await createBackup(formData);
      onSuccess?.();
    } catch (error) {
      console.error('Failed to create backup:', error);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="connection">Database Connection</Label>
        <Select
          value={formData.connectionId}
          onValueChange={(value) => setFormData({ ...formData, connectionId: value })}
        >
          <SelectTrigger>
            <SelectValue placeholder="Select a connection" />
          </SelectTrigger>
          <SelectContent>
            {connections?.map((conn) => (
              <SelectItem key={conn.id} value={conn.id}>
                {conn.name} ({conn.type})
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </div>

      <div className="space-y-2">
        <Label htmlFor="name">Backup Name</Label>
        <Input
          id="name"
          required
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, name: e.target.value })}
        />
      </div>

      <div className="space-y-2">
        <Label htmlFor="description">Description (Optional)</Label>
        <Input
          id="description"
          value={formData.description}
          onChange={(e) => setFormData({ ...formData, description: e.target.value })}
        />
      </div>

      <div className="flex items-center space-x-2">
        <Switch
          id="scheduled"
          checked={formData.scheduled}
          onCheckedChange={(checked) => setFormData({ ...formData, scheduled: checked })}
        />
        <Label htmlFor="scheduled">Schedule Backup</Label>
      </div>

      {formData.scheduled && (
        <div className="space-y-2">
          <Label htmlFor="cronExpression">Cron Expression</Label>
          <Input
            id="cronExpression"
            placeholder="*/5 * * * *"
            value={formData.cronExpression}
            onChange={(e) => setFormData({ ...formData, cronExpression: e.target.value })}
          />
          <p className="text-sm text-muted-foreground">
            Use cron expression format (e.g., &quot;0 0 * * *&quot; for daily at midnight)
          </p>
        </div>
      )}

      <div className="flex space-x-2">
        <Button type="submit" disabled={isSubmitting || !formData.connectionId}>
          {isSubmitting ? 'Creating Backup...' : 'Create Backup'}
        </Button>
        <Button type="button" variant="outline" onClick={onCancel}>
          Cancel
        </Button>
      </div>
    </form>
  );
}
