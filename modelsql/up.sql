CREATE TABLE messages (
                          id SERIAL PRIMARY KEY,
                          text TEXT NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                          processed BOOLEAN NOT NULL DEFAULT false
);


