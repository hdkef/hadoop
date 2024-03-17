CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE metadata (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    parent_path VARCHAR(250) NOT NULL,
    path VARCHAR(250) NOT NULL UNIQUE,
    m_type SMALLINT NOT NULL,
    i_node_id UUID,
    hash TEXT,
    all_block_ids UUID[]
);

CREATE INDEX idx_parent_path ON metadata(parent_path);
CREATE INDEX idx_path ON metadata USING HASH (path);
ALTER TABLE metadata ADD CONSTRAINT unique_i_node_id UNIQUE (i_node_id);