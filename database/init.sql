USE todo_db;

CREATE TABLE IF NOT EXISTS todos (
    id CHAR(36)
        CHARACTER SET ascii
        COLLATE ascii_bin
        NOT NULL
        PRIMARY KEY,
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