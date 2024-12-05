# Lesson 20: Migrations and Seeds

## Migrations

Database migrations are a way to manage and version changes to your database schema in a controlled and reproducible
manner.

- Track changes to your database schema over time.
- Synchronize the database schema with the application code.
- Avoid manual SQL execution, which can be error-prone.

Why do we need it?

- **Version Control for the Database**: Ensures that everyone on the team is working with the same schema.
- **Collaboration**: Developers can easily share changes and work on features without conflicts.
- **Deployment**: Automatically update the database schema during application deployment.
- **Rollback**: Easily revert changes if something goes wrong.

### Components

1. Migration Files:

   Text files (often written in SQL or a migration library syntax) that describe changes to the schema.

   Examples: Adding a column, creating a table, or updating constraints.

2. Migration Tools:

   Libraries or frameworks to apply migrations.

   Examples: go-migrate, Flyway, Liquibase, Django ORM, Rails ActiveRecord, Sequelize (Node.js).

3. Migration Workflow:

   **Create**: Write a migration file for the change.

   **Apply**: Run the migration to update the database.

   **Track**: Record applied migrations to prevent reapplying them.

### Flow

1. Start with a Migration Script:

   Create a migration using a tool or manually.

   ```sql
    -- migrations/20241205_add_birthdate_to_users.up.sql
    ALTER TABLE users ADD COLUMN birthdate DATE;
    ```
   ```sql
    -- migrations/20241205_add_birthdate_to_users.down.sql
    ALTER TABLE users DROP COLUMN birthdate;
    ```

2. Run the Migration:

   Use a tool command `migrate up`.

   Schema is updated in the database.

3. Review Changes:

   Confirm that the database schema matches expectations.

4. Rollback if Necessary:

   Use `migrate down` to revert changes.

### Best practices

- Use descriptive names for migration files (e.g., 20241205_add_birthdate_to_users).
- Always test migrations locally before applying to production. Especially rollback (down) migrations.
- Maintain backups before applying or rolling back migrations.
- Apply migrations in staging or development environments first.
- Include data migrations (e.g., populating new columns) when necessary.
- Keep migrations small and focused for easier debugging.

## Seeds

Database seeding is the process of populating a database with initial or default data that the application requires.

- Provide data to test or demo the application.
- Preload static or configuration data, such as roles, permissions, or categories.
- Quickly set up a working environment for development or testing.

Why do we need it?

- **Quick Setup**: Automatically populate essential data during setup.
- **Consistency**: Ensure all environments (development, staging, production) have the same initial data.
- **Testing**: Provide predictable datasets for testing.
- **Demos**: Create a populated database to showcase features.

### Components

1. Seed Data:

   The actual data to be inserted into the database.

   Example: User roles like "Admin," "User," and "Guest."

2. Seed Scripts:

   Code or files that define how the data is inserted.

   Examples: Go functions for Go, JSON files for Node.js.

3. Seeding Tools:

   Framework-specific tools or libraries that automate seeding.

   Examples: Sequelize seeds, Rails db:seed, or Django fixtures.

### Flow

1. Create Seed Data:

   Prepare data in the required format.

   ```go
   type user struct {
       id          int
       email       string
       role        string
       defaultName string
   
   }

   var users []*user
   ```

2. Write Seed Script:

   Create a script that inserts this data into the database.

3. Run the Seeder:

   Execute the script using the framework's command.

4. Verify the Data:

   Check the database to confirm that the data has been inserted correctly.

### Best practices

- **Organize Your Seed Data:** Split large datasets into smaller, modular scripts.
- **Version Control:** Store seed scripts in your code repository.
- **Environment Awareness:** Use different seeds for development, testing, and production.
- **Automation** Integrate seeding into setup scripts or CI/CD pipelines.
- **Reusability** Write reusable functions or utilities to generate seed data.
