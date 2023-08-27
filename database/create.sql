/*
 - Note:
 The balance is represented as an int because of floating point arithmetic.
 To cast the balance to a float, divide it by 100. Every user starts with 10000 cents (100.00).
 
 - Note:
 To apply the schema to the database, simply run: cat database/create.sql | sqlite3 database.sqlite
 */
-- Create schema
CREATE TABLE IF NOT EXISTS "users" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  cpf VARCHAR(11) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  balance INTEGER NOT NULL DEFAULT 10000,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transactions" (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  payer_id INTEGER NOT NULL,
  payee_id INTEGER NOT NULL,
  value INTEGER NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (payer_id) REFERENCES users(id),
  FOREIGN KEY (payee_id) REFERENCES users(id)
);

-- Add indexes
CREATE UNIQUE INDEX IF NOT EXISTS "users_email_unique" ON "users" ("email");

CREATE UNIQUE INDEX IF NOT EXISTS "users_cpf_unique" ON "users" ("cpf");