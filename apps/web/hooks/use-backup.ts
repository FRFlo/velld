import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getBackups, saveBackup } from "@/lib/api/backups";
import { useToast } from "@/hooks/use-toast";

export function useBackup() {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  const { data: backups, isLoading } = useQuery({
    queryKey: ["backups"],
    queryFn: getBackups,
  })

  const { mutate: addBackup, isPending: isAdding } = useMutation({
    mutationFn: async (connectionId: string) => {
      await saveBackup(connectionId);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["backups"],
      });
      toast({
        title: "Success",
        description: "Backup added successfully",
      });
    },
    onError: (error) => {
      toast({
        title: "Error",
        description: error.message || "Failed to add backup",
        variant: "destructive",
      });
    },
  })

  return {
    addBackup,
    isAdding,
    backups,
    isLoading,
  }
}