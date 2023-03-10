ALTER TABLE IF EXISTS "tb_payments" DROP CONSTRAINT IF EXISTS "tb_payments_ticket_id_fkey";
ALTER TABLE IF EXISTS "tb_payments" DROP CONSTRAINT IF EXISTS "tb_payments_user_id_fkey";
ALTER TABLE IF EXISTS "tb_halls" DROP CONSTRAINT IF EXISTS "tb_halls_theater_id_fkey";
ALTER TABLE IF EXISTS "tb_seats" DROP CONSTRAINT IF EXISTS "tb_seats_hall_id_fkey";
ALTER TABLE IF EXISTS "tb_showtimes" DROP CONSTRAINT IF EXISTS "tb_showtimes_hall_id_fkey";
ALTER TABLE IF EXISTS "tb_showtimes" DROP CONSTRAINT IF EXISTS "tb_showtimes_movie_id_fkey";
ALTER TABLE IF EXISTS "tb_tickets" DROP CONSTRAINT IF EXISTS "tb_tickets_seat_id_fkey";
ALTER TABLE IF EXISTS "tb_tickets" DROP CONSTRAINT IF EXISTS "tb_tickets_showtime_id_fkey";
ALTER TABLE IF EXISTS "tb_tickets" DROP CONSTRAINT IF EXISTS "tb_tickets_user_id_fkey";
ALTER TABLE IF EXISTS "tb_users" DROP CONSTRAINT IF EXISTS "tb_users_role_id_fkey";

DROP TABLE IF EXISTS tb_roles;
DROP TABLE IF EXISTS tb_users;
DROP TABLE IF EXISTS tb_theaters;
DROP TABLE IF EXISTS tb_halls;
DROP TABLE IF EXISTS tb_seats;
DROP TABLE IF EXISTS tb_movies;
DROP TABLE IF EXISTS tb_showtimes;
DROP TABLE IF EXISTS tb_tickets;
DROP TABLE IF EXISTS tb_payments;
