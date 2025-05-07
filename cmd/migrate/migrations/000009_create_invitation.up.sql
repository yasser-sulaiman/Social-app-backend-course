CREATE TABLE IF NOT EXISTS user_invitations (
    token BYTEA PRIMARY KEY,
    user_id BIGINT NOT NULL
)