import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getBackups, saveBackup } from "@/lib/api/backups";
import { useToast } from "@/hooks/use-toast";
import { useState } from 'react';

export function useBackup() {
  const queryClient = useQueryClient();
  const { toast } = useToast();

  const [page, setPage] = useState(1);
  const [limit] = useState(10);
  const [search, setSearch] = useState('');

  const { data, isLoading, error } = useQuery({
    queryKey: ['backups', { page, limit, search }],
    queryFn: () => getBackups({ page, limit, search }),
    placeholderData: (previousData) => previousData,
  });

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
    backups: data?.data,
    pagination: data?.pagination,
    isLoading,
    error,
    page,
    setPage,
    search,
    setSearch,
  }
}