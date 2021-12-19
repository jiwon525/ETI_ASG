# ETI_ASG
ETI Assignment: Creating a Taxi ride application
1. design consideration of my microservice
the passenger services such as passenger login, passenger register and editing passenger account details are not dependent on the driver functions or the trip functions. One functions that links them together would be the 'booking a ride' function, but it will still be functionable even if there are errors on one side. The three different main services are independent and loosely coupled. 
Each service can be independently upgraded or restarted as they are in separate .go files, and can be separately updated.\

2. architecture diagram
Architecture diagram done on a website called Lucid Chart : https://lucid.app/lucidchart/19649939-250b-4e35-8592-03ba5530744d/edit?viewport_loc=258%2C-23%2C2223%2C1242%2C0_0&invitationId=inv_f94163d0-f4f8-4af2-9618-6adc03b8be97

3. instructions for setting up and running the microservice
3. 1. Signing up as a new passenger / new driver 
3. 1. 1. Look at the menu and enter 2 if you want to sign up as a passenger, 4 if you want to sign up as a driver
3. 1. 2. Enter all the details requested 
