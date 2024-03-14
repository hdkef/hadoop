CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE i_nodes_blocks (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    blocks_id UUID NOT NULL,
    node_id VARCHAR(250) NOT NULL,
    blocks_index SMALLINT NOT NULL,
    size BIGINT NOT NULL,
    i_node_id UUID NOT NULL,
    CONSTRAINT fk_i_node_id FOREIGN KEY (i_node_id) REFERENCES metadata(i_node_id)
);