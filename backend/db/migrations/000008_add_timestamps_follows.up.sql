ALTER TABLE follows ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW();
ALTER TABLE follows ADD COLUMN deleted_at TIMESTAMPTZ;
