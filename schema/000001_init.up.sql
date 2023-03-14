CREATE TABLE IF NOT EXISTS roles
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS applicant_statuses
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS languages
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS language_levels
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

CREATE TABLE IF NOT EXISTS specializations
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS education_institutions
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS education_grades
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS vacancy_statuses
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS spheres
(
    id serial PRIMARY KEY,
    name character varying NOT NULL
);

CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    login character varying NOT NULL,
    password character varying NOT NULL,
    role_id integer REFERENCES roles (id),
    banned BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS employers
(
    user_id integer REFERENCES users (id),
    name character varying NOT NULL,
    phone character varying
);

CREATE TABLE IF NOT EXISTS applicant_info
(
    user_id integer REFERENCES users (id) NOT NULL,
    status_id integer REFERENCES applicant_statuses (id),
    name character varying NOT NULL,
    lastname character varying,
    phone character varying,
    birthday DATE,
    description character varying,
    male BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS experience
(
    user_id integer REFERENCES users (id) NOT NULL,
    organization character varying,
    start DATE,
    finish DATE,
    position character varying,
    duties character varying,
    skills character varying
);

CREATE TABLE IF NOT EXISTS education
(
    user_id integer REFERENCES users (id) NOT NULL,
    institution_id integer REFERENCES education_institutions (id),
    grade_id integer REFERENCES education_grades (id),
    faculty character varying,
    spezialization character varying,
    finish DATE
);

CREATE TABLE IF NOT EXISTS responses
(
    user_id integer REFERENCES users (id),
    vacancy_id character varying NOT NULL,
    status_id integer REFERENCES vacancy_statuses (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS employer_id_sphere_id
(
    user_id integer REFERENCES users (id) NOT NULL,
    sphere_id integer REFERENCES spheres (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS applicant_id_language_id
(
    user_id integer REFERENCES users (id) NOT NULL,
    language_id integer REFERENCES languages (id) NOT NULL,
    language_level integer REFERENCES language_levels (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS applicant_id_busyness_id
(
    user_id integer REFERENCES users (id) NOT NULL,
    busyness_id integer REFERENCES busyness (id)
);

CREATE TABLE IF NOT EXISTS applicant_id_schedule_id
(
    user_id integer REFERENCES users (id) NOT NULL,
    schedule_id integer REFERENCES schedule (id)
);

CREATE TABLE IF NOT EXISTS applicant_id_specialization_id
(
    user_id integer REFERENCES users (id) NOT NULL,
    specialization_id integer REFERENCES specializations (id)
);

CREATE TABLE IF NOT EXISTS employer_id_vacancy_id
(
    user_id integer REFERENCES users (id),
    vacancy_id character varying NOT NULL
);