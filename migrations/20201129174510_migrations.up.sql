CREATE TABLE employee (
    id INTEGER NOT NULL UNIQUE,
    name VARCHAR NOT NULL,
    secondName VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    hireDate VARCHAR NOT NULL,
    position VARCHAR NOT NULL,
    companyId INTEGER NOT NULL 
);