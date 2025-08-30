import { useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { Job } from "@/types/job";

export function useCreateJob() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: (job: Omit<Job, "id" | "created_at" | "updated_at">) =>
            api.post<Job>("/jobs", job),
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ["jobs"] });
        },
    });
}

