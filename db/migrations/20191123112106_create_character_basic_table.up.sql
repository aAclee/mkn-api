CREATE TABLE IF NOT EXISTS characters_basic (
  id SERIAL PRIMARY KEY,
  player_id INTEGER NOT NULL,
  campaign_id INTEGER,

  name text,
  family_name text,

  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);