# ETI_ASG
ETI Assignment: Creating a Taxi ride application
1. design consideration of my microservice
The passenger services such as passenger login, passenger register and editing passenger account details are not dependent on the driver functions or the trip functions. One functions that links them together would be the 'booking a ride' function, but it will still be functionable even if there are errors on one side. The three different main services are independent and loosely coupled. 
Each service can be independently upgraded or restarted as they are in separate .go files, and can be separately updated.
There is decomposition done by the business domain, as the passenger, driver and trip domains are separated. 
If I was utilising google cloud or AWS services, I would have used EC2 balancer to make sure that multiple servers are running and my service is horizontally and vertically scalable for higher resilience. Also, I would have my server on more than one region to further increase the resilience of my application. 

2. architecture diagram
Architecture diagram done on a website called Lucid Chart : https://lucid.app/lucidchart/19649939-250b-4e35-8592-03ba5530744d/edit?viewport_loc=258%2C-23%2C2223%2C1242%2C0_0&invitationId=inv_f94163d0-f4f8-4af2-9618-6adc03b8be97

3. instructions for setting up and running the microservice

 1. Signing up as a new passenger / new driver 
 1. 1. crom the console, enter 2 if you want to sign up as a passenger, 4 if you want to sign up as a driver
 1. 2. Enter all the details requested 

 2. Loging in as a passenger / new driver
 2. 1. from the console, enter 1 for passenger log in and 3 for driver log in. 
 2. 2. for passengers, enter your Username and your Password. the service will check if your input matches the data inside the database. 
        for drivers, enter your NRIC and your Password. the service will check if your input matches the data inside the database.

 3. After logging in as a Passenger
 3. 1. 1.To book a taxi ride, press 1 after you log in. 
 3. 1. 2.enter your current location postal code and your destination postal code. this will be saved in the database for the driver.
 3. 1. 3.Once inside the database and matched with a driver in the database, you will be shown the full details of the trip, like the destination postal codes, your username, your driver's name and your driver's ID. 
 3. 1. 4.Now you have to wait for the driver to accept and initiate the trip.

 3. 2. to retrieve past trips, press 2 after you log in. You will be able to see the details of your past trip.

 3. 3. 1. to edit account details other than the username, you can press 3 after you log in.
 3. 3. 2. input all the details prompted, and it will be reflected on the database. 

 3. 4. to log out, press 4.

4. After logging in as a Driver 
4. 1. as a driver, you are able to initiate the start and end of the trip assigned to you. 
4. 1. 1. to initiate trip, press 1.
4. 1. 2. to start the trip, press 1. if you have a trip assigned, it will be printed for you to accept or decline. 
4. 1. 3. to end the trip, after transportation, press 2. this will mark the end of the trip.

4. 2. as a driver, you can edit your account details, other than your NRIC.
4. 2. 1. to edit your account details, press 2. 
4. 2. 2. enter all the info prompted, and it will be reflected on the database. 

5. To exit the application, press 5 in the main menu.
