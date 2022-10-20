CREATE TABLE users
(
    api_key VARCHAR(255) NOT NULL,
    expiration_date VARCHAR(255) NOT NULL,

    PRIMARY KEY (api_key)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
