CREATE TABLE IF NOT EXISTS applicants
(
    id serial PRIMARY KEY,
    name character varying NOT NULL,
    lastname character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS specializations
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS employers
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS busyness
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS schedule
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS applicant_id_busyness_id
(
    applicant_id integer REFERENCES applicants (id),
    busyness_id integer REFERENCES busyness (id)
);

CREATE TABLE IF NOT EXISTS applicant_id_schedule_id
(
    applicant_id integer REFERENCES applicants (id),
    schedule_id integer REFERENCES schedule (id)
);

CREATE TABLE IF NOT EXISTS responses
(
    applicant_id integer REFERENCES applicants (id),
    vacancy_id character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS employer_id_vacancy_id
(
    employer_id integer REFERENCES employers (id),
    vacancy_id character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS applicant_id_specialization_id
(
    applicant_id integer REFERENCES applicants (id),
    specialization_id integer REFERENCES specializations (id)
);