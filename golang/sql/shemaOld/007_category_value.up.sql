CREATE SCHEMA category_value;
CREATE TABLE IF NOT EXISTS category_value
(
    id bigserial PRIMARY KEY,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    category_id bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    CONSTRAINT category_value_category_id_foreign FOREIGN KEY (category_id)
        REFERENCES category (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE RESTRICT
        NOT VALID
)