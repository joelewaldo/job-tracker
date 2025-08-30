import { Button } from "@/components/ui/button"

interface SidebarProps {
    currentView: string
    setView: (view: string) => void
    onAddJob: () => void
}

export function Sidebar({ currentView, setView, onAddJob }: SidebarProps) {
    return (
        <aside className="w-64 h-screen bg-muted p-4 flex flex-col gap-4">
            <h2 className="text-lg font-bold">Job Tracker</h2>
            <nav className="flex flex-col gap-2">
                <Button
                    variant={currentView === "jobs" ? "default" : "ghost"}
                    onClick={() => setView("jobs")}
                >
                    Jobs
                </Button>
                <Button
                    variant={currentView === "stats" ? "default" : "ghost"}
                    onClick={() => setView("stats")}
                >
                    Stats
                </Button>
            </nav>
            <div className="mt-auto">
                <Button className="w-full" onClick={onAddJob}>
                    + Add Job
                </Button>
            </div>
        </aside>
    )
}
