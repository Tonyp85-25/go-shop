CREATE TYPE user_role AS ENUM ('customer', 'admin');

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    public_id UUID NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    role user_role DEFAULT 'customer',
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_users_email ON users (email);

CREATE INDEX idx_users_deleted_at ON users (deleted_at);