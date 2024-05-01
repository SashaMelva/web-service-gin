-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.events DROP COLUMN date_time_end;
ALTER TABLE public.events DROP COLUMN date_time_start;
ALTER TABLE public.events DROP COLUMN date_time_send;

ALTER TABLE public.events ADD COLUMN date_time_end timestamp with time zone;
ALTER TABLE public.events ADD COLUMN date_time_start timestamp with time zone NOT NULL;
ALTER TABLE public.events ADD COLUMN date_time_send timestamp with time zone;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.events DROP COLUMN date_time_end;
ALTER TABLE public.events DROP COLUMN date_time_start;
ALTER TABLE public.events DROP COLUMN date_time_send;d;

ALTER TABLE public.events ADD COLUMN date_time_end time with time zone;
ALTER TABLE public.events ADD COLUMN date_time_start time with time zone NOT NULL;
ALTER TABLE public.events ADD COLUMN date_time_send time with time zone;
-- +goose StatementEnd
