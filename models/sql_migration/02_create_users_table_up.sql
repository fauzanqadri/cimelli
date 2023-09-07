-- cimelli_db.users definition

CREATE TABLE cimelli_db.users (
	id BIGINT NOT NULL,
	name varchar(255) NOT NULL,
	username varchar(255) NOT NULL,
	encrypted_password varchar(255) NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
DROP TABLE IF EXISTS cimelli_db.users;
CREATE UNIQUE INDEX users_id_IDX USING BTREE ON cimelli_db.users (id);
CREATE UNIQUE INDEX users_username_IDX USING BTREE ON cimelli_db.users (username);