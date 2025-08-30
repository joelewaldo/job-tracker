import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { Job } from "@/types/job"

interface JobsTableProps {
    jobs: Job[]
}

export function JobsTable({ jobs }: JobsTableProps) {
    return (
        <Table>
            <TableHeader>
                <TableRow>
                    <TableHead>Title</TableHead>
                    <TableHead>Company</TableHead>
                    <TableHead>Status</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {jobs.map((job) => (
                    <TableRow key={job.id}>
                        <TableCell>{job.position}</TableCell>
                        <TableCell>{job.company}</TableCell>
                        <TableCell>{job.status}</TableCell>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
    )
}
