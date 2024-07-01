# Server App

### Prerequisites
- Database MySQL
- Golang 1.20
- Import Database from file attached in root directory
  

## Description
This app using mix MVC and Clean Architeture concepts 

## Need to be improved
- In Create order usecase, to get latest product stock i still use locking table (row level locking (pesimistic locking) to avoid race condition). This would be better if you used a queuing system such as Redis Stream or Kafka
- User permission rules have not been implemented here, so that after logging in, anyone can manage customers 
