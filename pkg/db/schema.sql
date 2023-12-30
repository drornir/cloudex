CREATE TABLE "users" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"email"	text,
	UNIQUE("email")
);


CREATE TABLE "licenses" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"user" INTEGER NOT NULL,
	"product"	TEXT NOT NULL,
	"credentials"	TEXT NOT NULL,
	FOREIGN KEY("user") REFERENCES "users"("id")
);