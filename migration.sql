CREATE TABLE books (
    id INT GENERATED ALWAYS AS IDENTITY,
    author VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)