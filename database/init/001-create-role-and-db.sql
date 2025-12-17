DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'sreuser') THEN
    CREATE ROLE sreuser LOGIN PASSWORD 'srepassword';
  END IF;
END $$;

