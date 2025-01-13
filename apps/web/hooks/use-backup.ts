import { saveBackup } from "@/lib/api/backups";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "@/hooks/use-toast";

export function useBackup() {
  const queryClient = useQueryClient();

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
  }
}