import { useState } from "react"
import { Sidebar } from "@/components/Sidebar"
import { JobsTable } from "@/components/JobTable"
import { Job } from "@/types/job"
import './App.css'

function App() {
    const [view, setView] = useState("jobs")
    const [jobs, setJobs] = useState<Job[]>([
        {
            id: 1,
            position: "Software Engineer",
            company: "OpenAI",
            description: "Work on AI models and infrastructure",
            status: "applied",
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString(),
        },
        {
            id: 2,
            position: "DevOps Engineer",
            company: "Blue Origin",
            description: "Manage cloud infrastructure and deployment pipelines",
            status: "interview",
            created_at: new Date().toISOString(),
            updated_at: new Date().toISOString(),
        },
    ])
    const addJob = () => {
        const now = new Date().toISOString()

        const newJob: Job = {
            id: jobs.length + 1,
            position: "New Job",
            company: "Company X",
            description: "Job description goes here",
            status: "applied",
            created_at: now,
            updated_at: now,
        }

        setJobs([...jobs, newJob])
    }

    return (
        <div className="flex">
            <Sidebar currentView={view} setView={setView} onAddJob={addJob} />
            <main className="flex-1 p-6">
                {view === "jobs" && <JobsTable jobs={jobs} />}
                {view === "stats" && <p>Stats view coming soon 🚀</p>}
            </main>
        </div>
    )
}

export default App

