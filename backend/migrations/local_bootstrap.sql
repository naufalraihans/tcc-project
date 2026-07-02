CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS extensions;

CREATE TABLE IF NOT EXISTS auth.users (
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	email text,
	raw_user_meta_data jsonb DEFAULT '{}'::jsonb,
	created_at timestamptz DEFAULT now()
);

CREATE OR REPLACE FUNCTION auth.uid() RETURNS uuid
LANGUAGE sql STABLE AS $$ SELECT NULL::uuid $$;
