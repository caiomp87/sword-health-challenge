CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(50) NOT NULL,
    type VARCHAR(15) NOT NULL,
    email VARCHAR(50) NOT NULL,
    passwordHash VARCHAR(2000) NOT NULL,
    createdAt DATETIME NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tasks (
    id VARCHAR(36) NOT NULL, 
    name VARCHAR(50) NOT NULL, 
    summary VARCHAR(2500) NOT NULL,
    performed BOOLEAN NOT NULL DEFAULT false,
    createdAt DATETIME NOT NULL, 
    performedAt DATETIME NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
