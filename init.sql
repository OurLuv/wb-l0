
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS delivery_id_seq;

-- Table Definition
CREATE TABLE "public"."delivery" (
    "id" int4 NOT NULL DEFAULT nextval('delivery_id_seq'::regclass),
    "name" varchar,
    "phone" varchar,
    "zip" varchar,
    "city" varchar,
    "address" varchar,
    "region" varchar,
    "email" varchar,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS item_id_seq;

-- Table Definition
CREATE TABLE "public"."items" (
    "id" int4 NOT NULL DEFAULT nextval('item_id_seq'::regclass),
    "chrt_id" int4,
    "track_number" varchar,
    "price" int4,
    "rid" varchar,
    "name" varchar,
    "sale" int4,
    "size" varchar,
    "total_price" int4,
    "nm_id" int4,
    "brand" varchar,
    "status" int4,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."orders" (
    "order_uuid" uuid NOT NULL DEFAULT gen_random_uuid(),
    "track_number" varchar,
    "entry" varchar,
    "delivery_id" int4,
    "payment_id" int4,
    "locale" varchar,
    "internal_signature" varchar,
    "customer_id" varchar,
    "delivery_service" varchar,
    "shardkey" varchar,
    "sm_id" int4,
    "date_created" timestamp,
    "oof_shard" varchar,
    PRIMARY KEY ("order_uuid")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Table Definition
CREATE TABLE "public"."orders_items" (
    "order_uuid" uuid DEFAULT gen_random_uuid(),
    "item_id" int4
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS payment_id_seq;

-- Table Definition
CREATE TABLE "public"."payment" (
    "transaction" uuid DEFAULT gen_random_uuid(),
    "request_id" varchar,
    "currency" varchar,
    "provider" varchar,
    "amount" int4,
    "payment_dt" int4,
    "bank" varchar,
    "delivery_cost" int4,
    "goods_total" int4,
    "custom_fee" int4,
    "id" int4 NOT NULL DEFAULT nextval('payment_id_seq'::regclass),
    PRIMARY KEY ("id")
);

