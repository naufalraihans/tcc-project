CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS extensions;

CREATE TABLE IF NOT EXISTS auth.users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email text,
	encrypted_password text,
	raw_user_meta_data jsonb DEFAULT '{}'::jsonb,
	created_at timestamptz DEFAULT now()
);

ALTER TABLE auth.users ADD COLUMN IF NOT EXISTS encrypted_password text;
CREATE UNIQUE INDEX IF NOT EXISTS auth_users_email_key ON auth.users (email);

CREATE OR REPLACE FUNCTION auth.uid() RETURNS uuid
LANGUAGE sql STABLE AS $$ SELECT NULL::uuid $$;
