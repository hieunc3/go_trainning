USE bookstore;
DROP TABLE IF EXISTS books;
CREATE TABLE books(id INT PRIMARY KEY AUTO_INCREMENT,name VARCHAR(255),price Float,publication_year INT,author VARCHAR(255));
INSERT INTO books(name,price,publication_year,author) VALUES("To Kill a Mockingbird",103.5,1960,"Harper Lee")
INSERT INTO books(name,price,publication_year,author) VALUES("The Great Gatsby",112.5,1980,"F. Scott Fitzgerald")
INSERT INTO books(name,price,publication_year,author) VALUES("One Hundred Years of Solitude",143.5,1975,"Gabriel García Márquez")
INSERT INTO books(name,price,publication_year,author) VALUES("A Passage to India",97.5,1985,"E.M. Forster")
INSERT INTO books(name,price,publication_year,author) VALUES("Invisible Man",103.5,1952,"Ralph Ellison")