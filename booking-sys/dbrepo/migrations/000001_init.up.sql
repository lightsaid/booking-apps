CREATE TABLE "tb_users" (
  "id" bigserial PRIMARY KEY,
  "role_id" bigint NOT NULL,
  "phone_number" varchar NOT NULL,
  "password" varchar,
  "name" varchar NOT NULL,
  "avatar" varchar,
  "openid" varchar,
  "unionid" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_roles" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "code" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_theaters" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "location" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_halls" (
  "id" bigserial PRIMARY KEY,
  "theater_id" bigint NOT NULL,
  "name" varchar NOT NULL,
  "total_seats" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_seats" (
  "id" bigserial PRIMARY KEY,
  "hall_id" bigint NOT NULL,
  "col_number" int NOT NULL,
  "row_number" int NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_movies" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "release_date" timestamptz NOT NULL,
  "director" varchar NOT NULL,
  "poster" varchar NOT NULL,
  "duration" int NOT NULL,
  "genre" varchar,
  "star" varchar,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_showtimes" (
  "id" bigserial PRIMARY KEY,
  "movie_id" bigint NOT NULL,
  "hall_id" bigint NOT NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_tickets" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "showtime_id" bigint NOT NULL,
  "seat_id" bigint NOT NULL,
  "price" int NOT NULL,
  "booking_date" timestamptz,
  "payment_status" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "tb_payments" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint,
  "ticket_id" bigint,
  "NumberOfSeats" int NOT NULL,
  "payment_date" timestamptz NOT NULL,
  "payment_method" varchar NOT NULL,
  "payment_amount" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE UNIQUE INDEX ON "tb_users" ("phone_number");

CREATE UNIQUE INDEX ON "tb_roles" ("code");

CREATE INDEX ON "tb_halls" ("theater_id");

CREATE INDEX ON "tb_seats" ("hall_id");

CREATE INDEX ON "tb_showtimes" ("movie_id");

CREATE INDEX ON "tb_showtimes" ("hall_id");

CREATE INDEX ON "tb_tickets" ("user_id");

CREATE INDEX ON "tb_payments" ("user_id");

CREATE INDEX ON "tb_payments" ("ticket_id");

COMMENT ON COLUMN "tb_seats"."status" IS '状态: A、B、N, 分别是: 可用、被预订、损坏';

COMMENT ON COLUMN "tb_movies"."director" IS '导演';

COMMENT ON COLUMN "tb_movies"."poster" IS '海报/封面';

COMMENT ON COLUMN "tb_movies"."duration" IS '时长，单位: 分钟';

COMMENT ON COLUMN "tb_movies"."genre" IS '类型';

COMMENT ON COLUMN "tb_movies"."star" IS '主演';

COMMENT ON COLUMN "tb_showtimes"."start_time" IS '放映时间';

COMMENT ON COLUMN "tb_showtimes"."end_time" IS '结束时间';

COMMENT ON COLUMN "tb_tickets"."price" IS '单价，单位: 分';

COMMENT ON COLUMN "tb_tickets"."booking_date" IS '下订日期，被预订时设置时间';

COMMENT ON COLUMN "tb_tickets"."payment_status" IS '支付状态:Y/N';

COMMENT ON COLUMN "tb_payments"."NumberOfSeats" IS '电影票数';

COMMENT ON COLUMN "tb_payments"."payment_method" IS '支付方式';

COMMENT ON COLUMN "tb_payments"."payment_amount" IS '支付总额, 单位：分';

ALTER TABLE "tb_users" ADD FOREIGN KEY ("role_id") REFERENCES "tb_roles" ("id");

ALTER TABLE "tb_halls" ADD FOREIGN KEY ("theater_id") REFERENCES "tb_theaters" ("id");

ALTER TABLE "tb_seats" ADD FOREIGN KEY ("hall_id") REFERENCES "tb_halls" ("id");

ALTER TABLE "tb_showtimes" ADD FOREIGN KEY ("movie_id") REFERENCES "tb_movies" ("id");

ALTER TABLE "tb_showtimes" ADD FOREIGN KEY ("hall_id") REFERENCES "tb_halls" ("id");

ALTER TABLE "tb_tickets" ADD FOREIGN KEY ("user_id") REFERENCES "tb_users" ("id");

ALTER TABLE "tb_tickets" ADD FOREIGN KEY ("showtime_id") REFERENCES "tb_showtimes" ("id");

ALTER TABLE "tb_tickets" ADD FOREIGN KEY ("seat_id") REFERENCES "tb_seats" ("id");

ALTER TABLE "tb_payments" ADD FOREIGN KEY ("user_id") REFERENCES "tb_users" ("id");

ALTER TABLE "tb_payments" ADD FOREIGN KEY ("ticket_id") REFERENCES "tb_tickets" ("id");
