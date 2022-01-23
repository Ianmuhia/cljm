create table church_events
(
    id           bigserial
        primary key,
    created_at   timestamp with time zone,
    updated_at   timestamp with time zone,
    deleted_at   timestamp with time zone,
    organizer_id bigint
        constraint fk_church_events_organizer
            references users
            on update cascade on delete set null,
    cover_image  text,
    title        text,
    sub_title    text,
    content      text,
    scheduled_on timestamp with time zone
);

alter table church_events
    owner to wise;

create index idx_church_events_organizer_id
    on church_events (organizer_id);

create index idx_church_events_deleted_at
    on church_events (deleted_at);


create table church_jobs
(
    duty            text,
    church_event_id bigint,
    CONSTRAINT fk_church_event
        FOREIGN KEY (church_event_id)
            REFERENCES church_events (id) on update cascade on delete set null,
    id              bigserial,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    deleted_at      timestamp with time zone

);

alter table church_jobs
    owner to wise;

create index idx_church_jobs_deleted_at
    on church_jobs (deleted_at);


SELECT ce.*, cj.*
FROM church_events AS ce,
     church_jobs AS cj
WHERE ce.id = cj.church_event_id
  and ce.deleted_at IS NULL
  and cj.deleted_at IS NULL;

ALTER TABLE church_jobs
    ADD PRIMARY KEY (id);
