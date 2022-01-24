```sql
SELECT child.ID, child.UserName, parent.UserName as 'ParentUserName'
FROM Users as child
LEFT JOIN Users as parent ON (child.Parent = parent.ID);
```