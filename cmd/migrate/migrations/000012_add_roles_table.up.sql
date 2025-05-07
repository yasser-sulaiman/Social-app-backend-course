CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    LEVEL INT NOT NULL DEFAULT 0,
    description TEXT
);

INSERT INTO
    roles (name, LEVEL, description)
VALUES
    (
        'user',
        1,
        'Can create posts and comments'
    );

INSERT INTO
    roles (name, LEVEL, description)
VALUES
    (
        'moderator',
        3,
        'Can update other users posts'
    );

INSERT INTO
    roles (name, LEVEL, description)
VALUES
    (
        'admin',
        3,
        'Administrator role with all permissions'
    );