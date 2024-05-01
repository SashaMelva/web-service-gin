-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.events
(
     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
     title character varying COLLATE pg_catalog."default" NOT NULL,
     date_time_start time with time zone NOT NULL,
     date_time_end time with time zone,
     description text COLLATE pg_catalog."default",
     date_time_send time with time zone,
     CONSTRAINT events_pkey PRIMARY KEY (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
