ALTER TABLE blogs
ADD COLUMN slug VARCHAR(255) NOT NULL UNIQUE;

CREATE INDEX idx_blogs_created_at
ON blogs (created_at);