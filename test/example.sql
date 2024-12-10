CREATE DATABASE haproxydb (
        id INT NOT NULL PRIMARY KEY,
);
GO

CREATE USER haproxyuser
    WITH PASSWORD = 'password';
GO

USE haproxydb
    GRANT VIEW SERVER STATE TO haproxyuser;
GO