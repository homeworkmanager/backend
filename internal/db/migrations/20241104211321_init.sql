-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Groups (
      group_id BIGSERIAL PRIMARY KEY,
      group_name TEXT NOT NULL,
      course SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS Users (
     user_id BIGSERIAL PRIMARY KEY,
     name TEXT NOT NULL,
     surname TEXT,
     email TEXT UNIQUE,
     password TEXT NOT NULL,
     role int8 NOT NULL default 1 check ( role in (1, 2 ,3) ),
     group_id INTEGER,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

     CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id)
);


CREATE TABLE  IF NOT EXISTS Subjects (
     subject_id BIGSERIAL PRIMARY KEY,
     subject_name TEXT NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE  IF NOT EXISTS Class (
      class_id BIGSERIAL PRIMARY KEY,
      group_id INTEGER,
      subject_id INTEGER,
      date DATE NOT NULL,
      number_par INTEGER NOT NULL,
      location TEXT NOT NULL,

    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id)
);

CREATE TABLE IF NOT EXISTS Homework (
    homework_id BIGSERIAL PRIMARY KEY,
    group_id INTEGER,
    subject_id INTEGER,
    homework_text TEXT NOT NULL,
    due_date DATE NOT NULL,
    class_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id),
    CONSTRAINT class_fk FOREIGN KEY (class_id)REFERENCES Class(class_id)

);

CREATE TABLE IF NOT EXISTS HomeworkSubmissions (
   submissions_id BIGSERIAL PRIMARY KEY,
   user_id INTEGER,
   homework_id INTEGER,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

   CONSTRAINT user_fk FOREIGN KEY (user_id)REFERENCES Users(user_id),
   CONSTRAINT homework_fk FOREIGN KEY (homework_id)REFERENCES Homework(homework_id)

);

CREATE TABLE IF NOT EXISTS SubjectNotes (
    note_id BIGSERIAL PRIMARY KEY,
    subject_id INTEGER,
    group_id INTEGER,
    semester INTEGER NOT NULL,
    note_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    CONSTRAINT group_fk FOREIGN KEY (group_id)REFERENCES Groups(group_id),
    CONSTRAINT subject_fk FOREIGN KEY (subject_id)REFERENCES Subjects(subject_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS Groups CASCADE;
DROP TABLE IF EXISTS Subjects CASCADE;
DROP TABLE IF EXISTS Class CASCADE;
DROP TABLE IF EXISTS Homework CASCADE;
DROP TABLE IF EXISTS HomeworkSubmissions CASCADE;
DROP TABLE IF EXISTS SubjectNotes CASCADE;
-- +goose StatementEnd
