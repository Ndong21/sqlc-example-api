
CREATE TABLE "thread" (
  "id" SERIAL PRIMARY KEY,
  "topic" VARCHAR(36) NOT NULL,
  "created_at" TIMESTAMP DEFAULT now()
);


CREATE TABLE "message" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "thread_id" INT NOT NULL,
  "sender" VARCHAR(100) NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMP DEFAULT now(),
  FOREIGN KEY (thread_id) REFERENCES thread (id)
);

CREATE TABLE "orders" (
  "id" VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid()::varchar(36),
  "item" TEXT NOT NULL,
  "amount" TEXT NOT NULL,
  "number" TEXT NOT NULL
)

