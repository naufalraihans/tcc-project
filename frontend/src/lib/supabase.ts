import { createClient } from '@supabase/supabase-js';
import { env } from '$env/dynamic/public';

const url = env.PUBLIC_SUPABASE_URL || 'http://localhost:54321';
const key = env.PUBLIC_SUPABASE_ANON_KEY || 'anon-key-placeholder';

export const supabase = createClient(url, key);
