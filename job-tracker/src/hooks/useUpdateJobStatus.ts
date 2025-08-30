import { useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "@/lib/api";
import type { JobStatus } from "@/types/job";

export function useUpdateJobStatus() {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, status }: { id: number; status: JobStatus }) =>
            api.patch(`/jobs/${id}/status`, { status }),
        onSuccess: (_, { id }) => {
            queryClient.invalidateQueries({ queryKey: ["jobs"] });
            queryClient.invalidateQueries({ queryKey: ["job", id] });
        },
    });
}

