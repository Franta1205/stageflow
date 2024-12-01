CREATE TABLE organizations (
                               id VARCHAR(36) NOT NULL PRIMARY KEY,
                               name VARCHAR(255) NOT NULL,
                               created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP(3) NOT NULL
);