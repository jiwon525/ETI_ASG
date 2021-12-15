# ETI_ASG
ETI Assignment: Creating a Taxi ride application
1. design consideration of my microservice
the passenger services such as passenger login, passenger register and editing passenger account details are not dependent on the driver functions or the trip functions. One functions that links them together would be the 'booking a ride' function, but it will still be functionable even if there are errors on one side. The three different main services are independent and loosely coupled. 
Each service can be independently upgraded or restarted as they are in separate .go files, and can be separately updated.
architecture diagram

instructions for setting up and running the microservice