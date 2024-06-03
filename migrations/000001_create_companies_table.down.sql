CREATE TABLE IF NOT EXISTS companies(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    contact text NOT NULL, 
    address text NOT NULL,
    country text NOT NULL,
    socity_number text NOT NULL,
    code text NOT NULL,
    vat_number text NOT NULL,
    phone_number text NOT NULL
    create_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    email citext UNIQUE NOT NULL
)