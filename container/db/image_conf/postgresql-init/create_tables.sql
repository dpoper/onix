DO
$$
BEGIN
    ---------------------------------------------------------------------------
    -- ITEM TYPE
    ---------------------------------------------------------------------------
	IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='item_type')
	THEN
        CREATE SEQUENCE item_type_id_seq
            INCREMENT 1
            START 1
            MINVALUE 1
            MAXVALUE 9223372036854775807
            CACHE 1;

        ALTER SEQUENCE item_type_id_seq
            OWNER TO onix;

        CREATE TABLE item_type
        (
            id INTEGER NOT NULL DEFAULT nextval('item_type_id_seq'::regclass),
            name CHARACTER VARYING(200) COLLATE pg_catalog."default",
            description CHARACTER VARYING(500) COLLATE pg_catalog."default",
            custom boolean DEFAULT FALSE,
            CONSTRAINT item_type_id_pk PRIMARY KEY (id),
            CONSTRAINT item_type_name_uc UNIQUE (name)
        )
        WITH (
            OIDS = FALSE
        )
        TABLESPACE pg_default;

        ALTER TABLE item_type
            OWNER to onix;

        INSERT INTO item_type(name, description) VALUES ('INVENTORY', 'An Ansible inventory.');
        INSERT INTO item_type(name, description) VALUES ('HOST-GROUP', 'An Ansible host group.');
        INSERT INTO item_type(name, description) VALUES ('HOST', 'An Operating System Host.');
	END IF;

    ---------------------------------------------------------------------------
    -- ITEM
    ---------------------------------------------------------------------------
	IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='item')
	THEN
        CREATE SEQUENCE item_id_seq
            INCREMENT 1
            START 1
            MINVALUE 1
            MAXVALUE 9223372036854775807
            CACHE 1;

        ALTER SEQUENCE item_id_seq
            OWNER TO onix;

        CREATE TABLE item
        (
            id bigint NOT NULL DEFAULT nextval('item_id_seq'::regclass),
            name CHARACTER VARYING(200) COLLATE pg_catalog."default",
            description text COLLATE pg_catalog."default",
            deployed boolean,
            item_type_id INTEGER,
            meta json,
            version bigint NOT NULL DEFAULT 1,
            created timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP(6),
            updated timestamp(6) with time zone,
            tag CHARACTER VARYING(300) COLLATE pg_catalog."default",
            key CHARACTER VARYING(50) NOT NULL COLLATE pg_catalog."default",
            CONSTRAINT item_id_pk PRIMARY KEY (id),
            CONSTRAINT item_key_uc UNIQUE (key),
            CONSTRAINT item_item_type_id_fk FOREIGN KEY (item_type_id)
                REFERENCES item_type (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
        )
        WITH (
            OIDS = FALSE
        )
        TABLESPACE pg_default;

        ALTER TABLE item
            OWNER to onix;

        CREATE INDEX fki_item_item_type_id_fk
            ON item USING btree
            (item_type_id)
            TABLESPACE pg_default;
	END IF;

    ---------------------------------------------------------------------------
    -- LINK
    ---------------------------------------------------------------------------
    IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='link')
	THEN
        CREATE SEQUENCE link_id_seq
            INCREMENT 1
            START 1
            MINVALUE 1
            MAXVALUE 9223372036854775807
            CACHE 1;

        ALTER SEQUENCE link_id_seq
            OWNER TO onix;

        CREATE TABLE link
        (
            id bigint NOT NULL DEFAULT nextval('link_id_seq'::regclass),
            meta json,
            description text COLLATE pg_catalog."default",
            version bigint,
            created TIMESTAMP(6) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP(6),
            updated timestamp(6) WITH TIME ZONE,
            start_item_id bigint,
            end_item_id bigint,
            role CHARACTER VARYING(200) COLLATE pg_catalog."default" NOT NULL,
            key CHARACTER VARYING(50) COLLATE pg_catalog."default" NOT NULL,
            tag CHARACTER VARYING(300) COLLATE pg_catalog."default",
            CONSTRAINT link_id_pk PRIMARY KEY (id),
            CONSTRAINT link_key_uc UNIQUE (key),
            CONSTRAINT link_end_item_id_fk FOREIGN KEY (end_item_id)
                REFERENCES item (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION,
            CONSTRAINT link_start_item_id_fk FOREIGN KEY (start_item_id)
                REFERENCES item (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
        )
        WITH (
            OIDS = FALSE
        )
        TABLESPACE pg_default;

        ALTER TABLE link
            OWNER to onix;

        CREATE INDEX fki_link_end_item_id_fk
            ON link USING btree
            (end_item_id)
            TABLESPACE pg_default;

        CREATE INDEX fki_link_start_item_id_fk
            ON link USING btree
            (start_item_id)
            TABLESPACE pg_default;
    END IF;

    ---------------------------------------------------------------------------
    -- DIMENSION
    ---------------------------------------------------------------------------
    IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='dimension')
	THEN
        CREATE SEQUENCE dimension_id_seq
            INCREMENT 1
            START 1
            MINVALUE 1
            MAXVALUE 9223372036854775807
            CACHE 1;

        ALTER SEQUENCE dimension_id_seq
            OWNER TO onix;

        CREATE TABLE dimension
        (
            id bigint NOT NULL DEFAULT nextval('dimension_id_seq'::regclass),
            item_id bigint,
            key CHARACTER VARYING(50) NOT NULL,
            value CHARACTER VARYING(100) COLLATE pg_catalog."default" NOT NULL,
            CONSTRAINT dim_value_pkey PRIMARY KEY (id),
            CONSTRAINT dimension_item_id_fk FOREIGN KEY (item_id)
                REFERENCES item (id) MATCH SIMPLE
                ON UPDATE NO ACTION
                ON DELETE NO ACTION
        )
        WITH (
            OIDS = FALSE
        )
        TABLESPACE pg_default;

        ALTER TABLE dimension
            OWNER to onix;

        CREATE INDEX fki_dimension_item_id_fk
            ON dimension USING btree
            (item_id)
            TABLESPACE pg_default;
    END IF;

    ---------------------------------------------------------------------------
    -- ITEM AUDIT
    ---------------------------------------------------------------------------
    IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='item_audit')
    THEN
        CREATE TABLE item_audit
        (
            operation CHAR(1) NOT NULL,
            stamp TIMESTAMP NOT NULL,
            userid text NOT NULL,
            id bigint,
            name CHARACTER VARYING(200) COLLATE pg_catalog."default",
            description CHARACTER VARYING(500) COLLATE pg_catalog."default",
            deployed boolean,
            item_type_id INTEGER,
            meta json,
            version bigint,
            created TIMESTAMP(6) with time zone,
            updated TIMESTAMP(6) with time zone,
            tag CHARACTER VARYING(300) COLLATE pg_catalog."default",
            key CHARACTER VARYING(50) COLLATE pg_catalog."default"
        );

        ALTER TABLE item_audit
            OWNER to onix;

        CREATE OR REPLACE FUNCTION audit_item() RETURNS TRIGGER AS $item_audit$
        BEGIN
            IF (TG_OP = 'DELETE') THEN
                INSERT INTO item_audit SELECT 'D', now(), user, OLD.*;
                RETURN OLD;
            ELSIF (TG_OP = 'UPDATE') THEN
                INSERT INTO item_audit SELECT 'U', now(), user, NEW.*;
                RETURN NEW;
            ELSIF (TG_OP = 'INSERT') THEN
                INSERT INTO item_audit SELECT 'I', now(), user, NEW.*;
                RETURN NEW;
            END IF;
            RETURN NULL; -- result is ignored since this is an AFTER trigger
        END;
        $item_audit$ LANGUAGE plpgsql;

        CREATE TRIGGER item_audit
        AFTER INSERT OR UPDATE OR DELETE ON item
            FOR EACH ROW EXECUTE PROCEDURE audit_item();
    END IF;

    ---------------------------------------------------------------------------
    -- LINK AUDIT
    ---------------------------------------------------------------------------
    IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='link_audit')
    THEN
        CREATE TABLE link_audit
        (
            operation CHAR(1) NOT NULL,
            stamp TIMESTAMP NOT NULL,
            userid text NOT NULL,
            id bigint,
            meta json,
            description CHARACTER VARYING(500) COLLATE pg_catalog."default",
            version bigint,
            created TIMESTAMP(6) with time zone,
            updated TIMESTAMP(6) with time zone,
            start_item_id bigint,
            end_item_id bigint,
            role CHARACTER VARYING(200) COLLATE pg_catalog."default",
            key CHARACTER VARYING(50) COLLATE pg_catalog."default",
            tag CHARACTER VARYING(300) COLLATE pg_catalog."default"
        );

        ALTER TABLE link_audit
            OWNER to onix;

        CREATE OR REPLACE FUNCTION audit_link() RETURNS TRIGGER AS $link_audit$
        BEGIN
            IF (TG_OP = 'DELETE') THEN
                INSERT INTO link_audit SELECT 'D', now(), user, OLD.*;
                RETURN OLD;
            ELSIF (TG_OP = 'UPDATE') THEN
                INSERT INTO link_audit SELECT 'U', now(), user, NEW.*;
                RETURN NEW;
            ELSIF (TG_OP = 'INSERT') THEN
                INSERT INTO link_audit SELECT 'I', now(), user, NEW.*;
                RETURN NEW;
            END IF;
            RETURN NULL; -- result is ignored since this is an AFTER trigger
        END;
        $link_audit$ LANGUAGE plpgsql;

        CREATE TRIGGER link_audit
        AFTER INSERT OR UPDATE OR DELETE ON link
          FOR EACH ROW EXECUTE PROCEDURE audit_link();
    END IF;

    ---------------------------------------------------------------------------
    -- ITEM_TYPE AUDIT
    ---------------------------------------------------------------------------
    IF NOT EXISTS (SELECT relname FROM pg_class WHERE relname='item_type_audit')
    THEN
        CREATE TABLE item_type_audit
        (
            operation CHAR(1) NOT NULL,
            stamp TIMESTAMP NOT NULL,
            userid text NOT NULL,
            id INTEGER,
            name CHARACTER VARYING(200) COLLATE pg_catalog."default",
            description CHARACTER VARYING(500) COLLATE pg_catalog."default",
            custom boolean
        );

        ALTER TABLE item_type_audit
            OWNER to onix;

        CREATE OR REPLACE FUNCTION audit_item_type() RETURNS TRIGGER AS $item_type_audit$
        BEGIN
            IF (TG_OP = 'DELETE') THEN
                INSERT INTO item_type_audit SELECT 'D', now(), user, OLD.*;
                RETURN OLD;
            ELSIF (TG_OP = 'UPDATE') THEN
                INSERT INTO item_type_audit SELECT 'U', now(), user, NEW.*;
                RETURN NEW;
            ELSIF (TG_OP = 'INSERT') THEN
                INSERT INTO item_type_audit SELECT 'I', now(), user, NEW.*;
                RETURN NEW;
            END IF;
            RETURN NULL; -- result is ignored since this is an AFTER trigger
        END;
        $item_type_audit$ LANGUAGE plpgsql;

        CREATE TRIGGER item_type_audit
        AFTER INSERT OR UPDATE OR DELETE ON item_type
          FOR EACH ROW EXECUTE PROCEDURE audit_item_type();
    END IF;
END;
$$