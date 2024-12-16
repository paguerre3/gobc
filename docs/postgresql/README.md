
# PostgreSQL

Here’s a concise summary of PostgreSQL's main features, including **triggers**, presented in table format for clarity:

| **Feature**           | **Description**                                                                                   | **Example/Notes**                                                                                     |
|------------------------|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| **Open Source**        | PostgreSQL is free and open-source under the PostgreSQL License.                                 | No licensing costs, suitable for startups and enterprises.                                            |
| **ACID Compliance**    | Ensures reliable transactions with **Atomicity**, **Consistency**, **Isolation**, and **Durability**. | Transactions either fully complete or roll back if an error occurs.                                   |
| **Extensibility**      | Supports custom functions, operators, and data types.                                            | Add extensions like `PostGIS` for geographical data or `pg_cron` for job scheduling.                  |
| **Advanced Indexing**  | Provides indexing techniques like B-Tree, GIN, GiST, BRIN, and Hash for performance optimization. | Use **GIN** for full-text search or **BRIN** for large, sequential datasets.                          |
| **Replication**        | Supports **streaming replication**, logical replication, and cascading replication.               | Useful for high availability and horizontal scaling.                                                  |
| **JSON/JSONB Support** | Allows storage and querying of JSON data, with efficient indexing for JSONB.                     | Combine relational and NoSQL-style data storage.                                                      |
| **Full-Text Search**   | Built-in support for advanced full-text search with ranking and highlighting.                    | Index document data for efficient text search using **TSVector** and **TSQuery**.                     |
| **Partitioning**       | Enables horizontal partitioning for large tables, improving query performance.                   | Declarative partitioning by range or list, e.g., partition sales data by year.                        |
| **Triggers**           | Automatic execution of a function in response to certain events in a table.                     | Useful for enforcing constraints, logging changes, or maintaining audit trails.                       |
| **Foreign Data Wrappers (FDW)** | Access and query external databases or data sources as if they were local tables.                      | Query data from MySQL, MongoDB, or even files using FDWs like `mysql_fdw` or `file_fdw`.               |
| **Concurrency**        | Implements **MVCC** (Multi-Version Concurrency Control) for handling multiple transactions.       | Avoids locking issues, allowing concurrent read and write operations.                                 |
| **Stored Procedures**  | Supports **pl/pgsql** and other procedural languages for custom logic and data manipulation.      | Create complex business logic directly in the database layer.                                         |
| **Security**           | Offers robust security features, including role-based access control and row-level security.     | Define permissions at a granular level for multi-tenant applications.                                 |

---

### **Triggers in PostgreSQL**

| **Aspect**             | **Description**                                                                                   | **Example/Notes**                                                                                     |
|------------------------|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| **Definition**          | Triggers are database functions executed automatically before or after events like `INSERT`, `UPDATE`, or `DELETE`. | Helps maintain consistency or implement custom logic without application changes.                     |
| **Types**               | - **BEFORE Trigger:** Executes before the event. <br> - **AFTER Trigger:** Executes after the event. | Use **BEFORE** to modify data before insertion or **AFTER** for auditing or logging.                  |
| **Row vs. Statement**   | - **Row-Level Trigger:** Executes for each affected row. <br> - **Statement-Level Trigger:** Executes once per statement. | Use row-level triggers for per-row operations, e.g., updating related rows in another table.          |
| **Execution**           | Defined as a function in PL/pgSQL or other supported languages, linked to a specific table/event. | Example: A trigger that updates stock levels after an order is placed.                                |
| **Example Code**        | Create a trigger to log changes to a table:                                                      | ```sql<br>CREATE OR REPLACE FUNCTION log_changes() RETURNS TRIGGER AS $$ BEGIN INSERT INTO audit_log(table_name, operation, old_data, new_data) VALUES (TG_TABLE_NAME, TG_OP, row_to_json(OLD), row_to_json(NEW)); RETURN NEW; END; $$ LANGUAGE plpgsql;<br>CREATE TRIGGER after_update_log AFTER UPDATE ON employees FOR EACH ROW EXECUTE FUNCTION log_changes();``` |
