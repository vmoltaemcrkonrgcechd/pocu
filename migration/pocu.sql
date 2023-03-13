DROP DATABASE IF EXISTS pocu;

CREATE DATABASE pocu;

\connect pocu

SET CLIENT_ENCODING = 'UTF-8';

CREATE TABLE weapon(
                       weapon_id SMALLSERIAL NOT NULL,
                       name VARCHAR(40) NOT NULL,
                       attack REAL NOT NULL,
                       weight REAL NOT NULL,
                       PRIMARY KEY (weapon_id)
);

CREATE TABLE armor(
                      armor_id SMALLSERIAL NOT NULL,
                      name VARCHAR(40) NOT NULL,
                      protection REAL NOT NULL,
                      weight REAL NOT NULL,
                      PRIMARY KEY (armor_id)
);

CREATE TABLE class(
                      class_id SMALLSERIAL NOT NULL,
                      name VARCHAR(40) NOT NULL,
                      health REAL NOT NULL,
                      attack_speed REAL NOT NULL,
                      stamina REAL NOT NULL,
                      PRIMARY KEY (class_id)
);

CREATE TABLE combat_unit(
                            combat_unit_id SMALLINT NOT NULL,
                            name VARCHAR(40) NOT NULL,
                            attack REAL NOT NULL DEFAULT 2.5,
                            protection REAL NOT NULL DEFAULT 2.5,
                            health REAL NOT NULL DEFAULT 25,
                            attack_speed REAL NOT NULL DEFAULT 0.75,
                            stamina REAL NOT NULL DEFAULT 25,
                            class_id SMALLINT,
                            armor_id SMALLINT,
                            weapon_id SMALLINT,
                            PRIMARY KEY (combat_unit_id),
                            FOREIGN KEY (class_id) REFERENCES class (class_id)
                                ON DELETE SET NULL,
                            FOREIGN KEY (armor_id) REFERENCES armor (armor_id)
                                ON DELETE SET NULL,
                            FOREIGN KEY (weapon_id) REFERENCES weapon (weapon_id)
                                ON DELETE SET NULL
);

CREATE TABLE enemy(
                      enemy_id SMALLINT NOT NULL,
                      name VARCHAR(40) NOT NULL,
                      attack REAL NOT NULL,
                      protection REAL NOT NULL,
                      health REAL NOT NULL,
                      attack_speed REAL NOT NULL,
                      stamina REAL NOT NULL
);

INSERT INTO weapon (name, attack, weight)
VALUES
    ('Двуручник тёмной луны', 80.5, 10.0),
    ('Полуторный меч', 138.5, 9.0),
    ('Редувия', 40.0, 2.5);
