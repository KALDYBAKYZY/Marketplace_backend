CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          name VARCHAR(255) UNIQUE NOT NULL,
                          price DECIMAL(12,2) NOT NULL,
                          stock INT NOT NULL DEFAULT 0,
                          category_id BIGINT NOT NULL,

                          CONSTRAINT fk_products_category
                              FOREIGN KEY (category_id)
                                  REFERENCES categories(id)
                                  ON DELETE CASCADE
);