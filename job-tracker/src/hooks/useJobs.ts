import { useQuery } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { Job, JobStatus } from "@/types/job";

export function useJobs(status?: JobStatus) {
    return useQuery({
        queryKey: ["jobs", status],
        queryFn: () => {
            const query = status ? `?status=${status}` : "";
            return api.get<Job[]>(`/jobs${query}`);
        },
    });
}

