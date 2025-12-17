DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'sreuser') THEN
    CREATE ROLE sreuser LOGIN PASSWORD 'srepassword';
  END IF;
END $$;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'sredb') THEN
    CREATE DATABASE sredb OWNER sreuser;
  END IF;
END $$;
