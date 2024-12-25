import { useMutation } from '@tanstack/react-query';
import { login, register } from '@/lib/api/auth';
import { useRouter } from 'next/navigation';
import { useToast } from '@/hooks/use-toast';

export function useLogin() {
  const router = useRouter();
  const { toast } = useToast();

  return useMutation({
    mutationFn: login,
    onSuccess: () => {
      toast({
        title: "Success",
        description: "Logged in successfully",
      });
      router.push("/dashboard");
    },
    onError: () => {
      toast({
        title: "Error",
        description: "Failed to login",
        variant: "destructive",
      });
    },
  });
}

export function useRegister() {
  const router = useRouter();
  const { toast } = useToast();

  return useMutation({
    mutationFn: register,
    onSuccess: () => {
      toast({
        title: "Success",
        description: "Registered successfully",
      });
      router.push("/login");
    },
    onError: () => {
      toast({
        title: "Error",
        description: "Failed to register",
        variant: "destructive",
      });
    },
  });
}
