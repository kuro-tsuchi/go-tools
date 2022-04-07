package test

const SqlMock = `CREATE TABLE person_info (
	id BIGINT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT 'primary id',
	name VARCHAR(30) NOT NULL DEFAULT 'default_name',
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	age INT(11) unsigned NULL,
	sex VARCHAR(2) NULL
	);`