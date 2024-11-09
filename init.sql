BEGIN;

CREATE TABLE IF NOT EXISTS public."users"
(
    "id"        SERIAL      NOT NULL,
    "name"      TEXT        NOT NULL,
    "email"     TEXT        NOT NULL,
    "modified"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created"   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS public."documents"
(
    "id"        SERIAL      NOT NULL,
    "name"      TEXT        NOT NULL,
    "mime"      TEXT        NOT NULL,
    "is_file"   BOOLEAN     NOT NULL,
    "is_public" BOOLEAN     NOT NULL,
    "modified"  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created"   TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS public."users2documents"
(
    "user_id"     INTEGER NOT NULL REFERENCES public."users" ("id")     ON DELETE CASCADE ON UPDATE NO ACTION,
    "document_id" INTEGER NOT NULL REFERENCES public."documents" ("id") ON DELETE CASCADE ON UPDATE NO ACTION,
    PRIMARY KEY ("user_id", "document_id")
);

COMMIT;
