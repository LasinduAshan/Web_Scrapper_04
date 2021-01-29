drop database scrapper2;
create database scrapper2;
use scrapper2;




CREATE TABLE car(
	id int auto_increment,
	district VARCHAR(225) NOT NULL,
	category VARCHAR(225) NOT NULL,
	title VARCHAR(50) NOT NULL,
	price VARCHAR(50),
	url VARCHAR(225) NOT NULL,
	postedOn VARCHAR(225),
	forSaleBy VARCHAR(225),
	meta TEXT NOT NULL,
	CONSTRAINT PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=latin1;