-- CreateTable
CREATE TABLE users (
    user_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    type TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CreateTable
CREATE TABLE functions (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    creator_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    language TEXT NOT NULL,
    timeout INTEGER NOT NULL,
    output_type TEXT NOT NULL DEFAULT 'json',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT functions_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES users (user_id) ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE function_outputs (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    function_id INTEGER NOT NULL,
    output TEXT NOT NULL,
    cost INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT function_outputs_function_id_fkey FOREIGN KEY (function_id) REFERENCES functions (id) ON DELETE RESTRICT ON UPDATE CASCADE

);

-- CreateIndex
CREATE UNIQUE INDEX users_email_key ON users(email);
