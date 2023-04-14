CREATE DATABASE banking;
USE banking;

CREATE TABLE customers (
    customer_id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    date_of_birth date NOT NULL,
    city varchar(100) NOT NULL,
    zipcode varchar(10) NOT NULL,
    status tinyint(1) NOT NULL DEFAULT 1,
    PRIMARY KEY(customer_id)
) ENGINE=InnoDB AUTO_INCREMENT=2006 CHARSET=latin1;

INSERT INTO customers VALUES 
    (2000, 'Evandro', '1992-02-26', 'União da Vitória', '84605655', 1),
    (2001, 'Nick', '1971-05-01', 'Delhi', '150022', 1),
    (2002, 'John', '1977-02-03', 'New York', '225211', 1),
    (2003, 'Anna', '1980-03-10', 'Detroid', '8257587', 0),
    (2004, 'Eddy', '2010-08-15', 'Rio de Janeiro', '785552', 1),
    (2005, 'Samy', '2002-11-24', 'Dallas', '205866', 0);