# cetecMehulTask

#For testing please change the username and password for connection of mysql database in the code

Task 1:
make a REST endpoint (GET)
/person/{person_id}/info
GET /person/1/info  returns:
{
"name": "mike",
"phone_number": "444-444-4444",
"city" : "Austin",
"state" : "TX",
"street1": "213 South 1st St",
"street2": "", 
"zip_code": "78704",
}

Task 2:
make a REST endpoint (POST)  / Create
/person/create
POST /person/create  accepts body: 
{
"name": "YOURNAME",
"phone_number": "123-456-7890",
"city" : "Sacramento",
"state" : "CA",
"street1": "112 Main St",
"street2": "Apt 12", 
"zip_code": "12345",
}
returns 200


Person
create table person(id int not null key auto_increment,  name varchar(255), age int);
insert into person(id, name, age) values (1, "mike", 31), (2, "John", 20), (3, "Joseph", 20);
You can assume the following data in person:
Name, age, id
mike , 31, 1
John, 45, 2
Joseph, 20, 3


phone
create table phone(id int not null key auto_increment,  number varchar(255), person_id int);
insert into phone(id, person_id, number) values (1,1, "444-444-4444"), (8,2, "123-444-7777"), (3,3, "445-222-1234");
You can assume the following data in phone:
person_id, id, number
1,1, 444-444-4444
2,8, 123-444-7777
3,3, 445-222-1234


address
create table address(id int not null key auto_increment,  city varchar(255), state varchar(255), street1 varchar(255), street2 varchar(255), zip_code varchar(255));
insert into address(id ,  city , state , street1 , street2 , zip_code ) values (1,"Eugene", "OR", "111 Main St", "", "98765"), (2, "Sacramento", "CA", "432 First St", "Apt 1", "22221"), (3, "Austin", "TX", "213 South 1st St", "", "78704");
You can assume the following data in address:
id, city, state, street1, street2, zip_code
1,Eugene, OR, "111 Main St", "", 98765
2, Sacramento, CA, "432 First St", "Apt 1", 22221
3, Austin, TX, "213 South 1st St", "", 78704