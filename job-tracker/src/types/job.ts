export type JobStatus =
    | "applied"
    | "interview"
    | "offer"
    | "rejected"
    | "accepted"
    | "archived";

export interface Job {
    id: number;
    company: string;
    position: string;
    description: string;
    status: JobStatus;
    created_at: string;
    updated_at: string;
}
