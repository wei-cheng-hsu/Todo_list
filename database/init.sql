USE todo_db;

CREATE TABLE IF NOT EXISTS todos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,

    title VARCHAR(50) NOT NULL,

    description VARCHAR(1000),

    due_date DATE NOT NULL,

    setup_date DATETIME
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP,

    status ENUM('undo', 'finish', 'delay')
        NOT NULL
        DEFAULT 'undo',

    updated_at DATETIME
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);