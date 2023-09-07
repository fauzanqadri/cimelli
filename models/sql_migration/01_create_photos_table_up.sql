-- cimelli_db.photos definition

CREATE TABLE cimelli_db.photos (
	id BIGINT(20) NOT NULL,
	name varchar(100) NOT NULL,
	file_location varchar(255) NOT NULL,
	url_path varchar(255) NOT NULL,
	file_size INT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT photos_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
CREATE UNIQUE INDEX photos_name_IDX USING BTREE ON cimelli_db.photos (name);