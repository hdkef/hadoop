CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    is_committed BOOLEAN NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    lease_time_in_sec INTEGER NOT NULL,
    protobuf_bytes BYTEA NOT NULL
);

ALTER TABLE transactions
ADD COLUMN IF NOT EXISTS rollback_tries SMALLINT DEFAULT 0,
ADD COLUMN IF NOT EXISTS is_rolledback BOOLEAN DEFAULT false;