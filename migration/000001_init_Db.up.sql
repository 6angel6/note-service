CREATE TABLE users (
                       user_id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                       username VARCHAR(20) NOT NULL UNIQUE,
                       password VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notes (
                       note_id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
                       user_id INTEGER NOT NULL,
                       content TEXT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (user_id) REFERENCES users (user_id)
);


INSERT INTO users (username, password)
VALUES
    ('user_test', '123456'),
    ('bot_test', 'qwerty');


INSERT INTO notes (user_id, content)
VALUES
    (1, 'First note for user_test.'),
    (1, 'Second note for user_test.'),
    (2, 'First note for bot_test.'),
    (2, 'Second note for bot_test.');