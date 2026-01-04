CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          price BIGINT NOT NULL,
                          created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_products_name ON products(name);
