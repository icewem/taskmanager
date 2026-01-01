CREATE TABLE IF NOT EXISTS tasks (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     job_name  TEXT NOT NULL,
                                     start_at  TEXT,
                                     stop_at   TEXT,
                                     is_close  BOOLEAN NOT NULL,
                                     priority  TEXT,
                                     tags      TEXT -- JSON
);
