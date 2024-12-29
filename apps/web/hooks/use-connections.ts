import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { saveConnection, testConnection, getConnections, getConnectionStats } from '@/lib/api/connections';
import { DatabaseConnection } from '@/types/connection';
import { useToast } from '@/hooks/use-toast';

export function useConnections() {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  const { data: connections, isLoading } = useQuery({
    queryKey: ['connections'],
    queryFn: getConnections,
  });

  const { mutate: addConnection, isPending: isAdding } = useMutation({
    mutationFn: async (connection: DatabaseConnection) => {
      await testConnection(connection);
      await saveConnection(connection);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['connections'] });
      toast({
        title: "Success",
        description: "Connection added successfully",
      });
    },
    onError: (error) => {
      toast({
        title: "Error",
        description: error.message || "Failed to add connection",
        variant: "destructive",
      });
    },
  });

  const { data: stats } = useQuery({
    queryKey: ['connection-stats'],
    queryFn: getConnectionStats,
  });

  return {
    connections,
    isLoading,
    addConnection,
    isAdding,
    stats,
  };
}
