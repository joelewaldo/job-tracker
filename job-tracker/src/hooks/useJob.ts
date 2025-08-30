import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { Job } from "@/types/job";

export function useJob(id: number) {
    return useQuery({
        queryKey: ["job", id],
        queryFn: () => api.get<Job>(`/jobs/${id}`),
        enabled: !!id,
    });
}

