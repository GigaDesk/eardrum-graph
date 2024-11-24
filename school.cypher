//Create a constraint requiring School nodes to have unique pk properties
CREATE CONSTRAINT school_pk
FOR (school:School) REQUIRE school.pk IS UNIQUE;

//Create a constraint requiring School nodes to have unique phonenumber properties
CREATE CONSTRAINT school_phonenumber
FOR (school:School) REQUIRE school.phonenumber IS UNIQUE;