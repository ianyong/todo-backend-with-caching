-- +migrate Up
CREATE TABLE todos (
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR NOT NULL,
    description  VARCHAR NOT NULL,
    due_date     TIMESTAMP WITH TIME ZONE NOT NULL,
    is_completed BOOLEAN DEFAULT false
);

-- +migrate Down
DROP TABLE todos;
