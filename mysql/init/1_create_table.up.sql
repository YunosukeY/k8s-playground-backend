CREATE DATABASE IF NOT EXISTS sample_db;
CREATE TABLE IF NOT EXISTS sample_db.todos(
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  content VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);