-- +goose Up
-- +goose StatementBegin
CREATE TABLE faqs (
  id SERIAL PRIMARY KEY,
  category_id INTEGER NOT NULL,
  store_id INTEGER,
  is_global BOOLEAN DEFAULT FALSE,
  question VARCHAR(500) NOT NULL,
  answer TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_faqs_category FOREIGN KEY (category_id) REFERENCES faq_categories(id) ON DELETE CASCADE,
  CONSTRAINT fk_faqs_store FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
);

CREATE INDEX idx_faqs_category_id ON faqs(category_id);
CREATE INDEX idx_faqs_store_id ON faqs(store_id);
CREATE INDEX idx_faqs_is_global ON faqs(is_global);
CREATE INDEX idx_faqs_deleted_at ON faqs(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS faqs;
-- +goose StatementEnd
