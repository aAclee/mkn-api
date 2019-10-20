CREATE TABLE IF NOT EXISTS campaigns (
  id SERIAL PRIMARY KEY,
  name text NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);