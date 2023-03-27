CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost'; 
-- Important: Make sure to swap 'pass' with a password of your own choosing. 
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';

-- Test the new user (password is pass)
-- mysql -D snippetbox -u web -p