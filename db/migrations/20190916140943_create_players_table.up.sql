CREATE TABLE IF NOT EXISTS players (
  id SERIAL PRIMARY KEY,
  uuid UUID NOT NULL UNIQUE,
  email text NOT NULL UNIQUE,
  admin BOOLEAN DEFAULT false,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);