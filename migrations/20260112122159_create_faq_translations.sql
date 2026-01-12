-- +goose Up
-- +goose StatementBegin
CREATE TABLE faq_translations (
  id SERIAL PRIMARY KEY,
  faq_id INTEGER NOT NULL,
  language VARCHAR(10) NOT NULL,
  question VARCHAR(500) NOT NULL,
  answer TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_translations_faq FOREIGN KEY (faq_id) REFERENCES faqs(id) ON DELETE CASCADE,
  CONSTRAINT uq_faq_language UNIQUE (faq_id, language)
);

CREATE INDEX idx_faq_translations_faq_id ON faq_translations(faq_id);
CREATE INDEX idx_faq_translations_language ON faq_translations(language);
CREATE INDEX idx_faq_translations_deleted_at ON faq_translations(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS faq_translations;
-- +goose StatementEnd
