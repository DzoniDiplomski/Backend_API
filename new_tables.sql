-- Table: public.kalkulacija

-- DROP TABLE IF EXISTS public.kalkulacija;

CREATE TABLE IF NOT EXISTS public.kalkulacija
(
    id integer NOT NULL DEFAULT nextval('kalkulacija_id_seq'::regclass),
    "createdAt" date NOT NULL DEFAULT CURRENT_DATE,
    CONSTRAINT kalkulacija_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.kalkulacija
    OWNER to postgres;

-- Table: public.sadrzi_6

-- DROP TABLE IF EXISTS public.sadrzi_6;

CREATE TABLE IF NOT EXISTS public.sadrzi_6
(
    id_kalkulacije bigint NOT NULL,
    id_stavke bigint NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.sadrzi_6
    OWNER to postgres;

-- Table: public.stavka_kalkulacije

-- DROP TABLE IF EXISTS public.stavka_kalkulacije;

CREATE TABLE IF NOT EXISTS public.stavka_kalkulacije
(
    id integer NOT NULL DEFAULT nextval('stavka_kalkulacije_id_seq'::regclass),
    sif bigint,
    neto_cena double precision,
    marza double precision,
    kolicina integer,
    pdv_stopa integer,
    CONSTRAINT stavka_kalkulacije_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.stavka_kalkulacije
    OWNER to postgres;