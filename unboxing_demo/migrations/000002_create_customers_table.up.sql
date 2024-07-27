CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,  // Added created_at field
    name TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    address TEXT,
    version INT NOT NULL DEFAULT 1
);
