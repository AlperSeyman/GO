Project Structure Guide

1. cmd/ (The Ignition)

    Purpose: Contains the main entry points for the application.

    Contents: Usually has a subfolder like api/ containing main.go.

    Note: Keep this folder small. It should only initialize and start the app, not contain business logic.

2. internal/ (The Private Heart)

    Purpose: Contains code that is private to this project. Go prevents other external projects from importing anything inside this folder.

    Contents: This is where the "heavy lifting" happens.

3. internal/api/ (The Traffic Control)

    routers/: The "Map" of the application. It connects URL paths (like /users) to specific Handlers.

    middlewares/: The "Security Guard." Functions that run before the main logic to handle authentication, logging, or error recovery.

    handlers/: The "Receptionist." They receive HTTP requests, validate input, call the necessary logic, and send back HTTP responses.

4. internal/model/ (The Blueprints)

    Purpose: Defines the data structures (Structs) used across the application.

    Contents: Definitions like type User struct {...} or type Product struct {...}.

5. internal/repository/ (The Data Vault)

    Purpose: Handles all communication with the database.

    Contents: sqlconnect/ usually holds the logic for connecting to and querying the database (SQL).

    Note: By isolating SQL here, you can change your database type without touching your API or Model code.

6. pkg/ (The Shared Toolbox)

    Purpose: Contains helper code that is public. It can be imported and used by other projects.

    Contents: utils/ contains generic tools like string formatters, date helpers, or encryption functions that aren't specific to "School API" logic.