-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  user_type VARCHAR(20) NOT NULL CHECK (user_type IN ('customer', 'merchant', 'admin')),
  store_id INTEGER,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE SET NULL,
  CHECK (user_type = 'merchant' OR store_id IS NULL),
  CHECK (user_type != 'merchant' OR store_id IS NOT NULL)
);
CREATE INDEX idx_users_store_id ON users(store_id);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
INSERT INTO users (email, password, user_type, created_at, updated_at) 
VALUES ('admin@yamm.com', '$2a$14$wP9Py14jT7AlH.qL2y6CXenip.7b2xXrHJWzb8r6dpMLKRkL7Wl8S', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
