CREATE TABLE webinars (
                          id VARCHAR(36) NOT NULL PRIMARY KEY,
                          title VARCHAR(255) NOT NULL,
                          organisation_id VARCHAR(36) NOT NULL,
                          FOREIGN KEY (organisation_id) REFERENCES organisations(id) ON DELETE CASCADE,
                          scheduled_at TIMESTAMP(3),
                          created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP(3) NOT NULL
);