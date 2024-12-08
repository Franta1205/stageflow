create TABLE participants (
                              id VARCHAR(36) NOT NULL PRIMARY KEY,
                              email VARCHAR(255) NOT NULL,
                              first_name VARCHAR(255) NOT NULL,
                              last_name VARCHAR(255) NOT NULL,
                              webinar_id VARCHAR(36) NOT NULL,
                              FOREIGN KEY (webinar_id) REFERENCES organisations(id),
                              created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP(3) NOT NULL
);