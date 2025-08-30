CREATE TABLE IF NOT EXISTS jobs (
    id SERIAL PRIMARY KEY,
    company TEXT NOT NULL,
    position TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL CHECK (status IN ('applied', 'interview', 'offer', 'rejected', 'accepted', 'archived')) DEFAULT 'applied',
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

