CREATE TABLE user_organisations (
                                    user_id VARCHAR(36) NOT NULL,
                                    organisation_id VARCHAR(36) NOT NULL,
                                    role VARCHAR(50) NOT NULL DEFAULT 'owner',
                                    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP(3) NOT NULL,
                                    PRIMARY KEY (user_id, organisation_id),
                                    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                                    FOREIGN KEY (organisation_id) REFERENCES organisations(id) ON DELETE CASCADE
);