CREATE TABLE users
(
    api_key VARCHAR(255) NOT NULL,
    expiration_date VARCHAR(255) NOT NULL,

    PRIMARY KEY (api_key)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

CREATE TABLE user_requests
(
    api_key VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    headers JSON NOT NULL,
    CHECK (JSON_VALID(headers))
) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

