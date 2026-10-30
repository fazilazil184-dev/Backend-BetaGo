DROP INDEX IF EXISTS idx_blogs_created_at;

ALTER TABLE blogs
DROP COLUMN slug;