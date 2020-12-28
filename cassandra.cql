CREATE KEYSPACE IF NOT EXISTS em WITH REPLICATION = { 'class': 'SimpleStrategy',
'replication_factor': 1 };

CREATE TABLE IF NOT EXISTS em.messages (
  id UUID,
  timestamp TIMESTAMP,
  email TEXT,
  title TEXT,
  content TEXT,
  magic_number INT,
  PRIMARY KEY(id)
);